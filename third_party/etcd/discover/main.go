package main

import (
	"context"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type ServiceDiscovery struct {
	cli      *clientv3.Client
	services map[string]string
	lock     sync.Mutex
}

func NewServiceDiscovery(endpoints []string) *ServiceDiscovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceDiscovery{
		cli:      cli,
		services: make(map[string]string),
	}
}

func (service *ServiceDiscovery) WatchServices(prefix string) error {
	resp, err := service.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, kv := range resp.Kvs {
		service.SetService(string(kv.Key), string(kv.Value))
	}

	// watch prefix
	go service.watcher(prefix)
	return nil
}

func (service *ServiceDiscovery) watcher(prefix string) {
	watch := service.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Println("Watching prefix: ", prefix)
	for resp := range watch {
		for _, e := range resp.Events {
			switch e.Type {
			case 0:
				service.SetService(string(e.Kv.Key), string(e.Kv.Value))
			case 1:
				service.DeleteService(string(e.Kv.Key))
			}
		}
	}
}

func (service *ServiceDiscovery) SetService(key, val string) {
	service.lock.Lock()
	defer service.lock.Unlock()
	service.services[key] = val
	log.Println("PUT Key: ", key, " Value: ", val)
}

func (service *ServiceDiscovery) DeleteService(key string) {
	service.lock.Lock()
	defer service.lock.Unlock()
	delete(service.services, key)
	log.Println("DELETE Key: ", key)
}

func (service *ServiceDiscovery) GetServices() []string {
	service.lock.Lock()
	defer service.lock.Unlock()
	addrs := make([]string, 0)
	for _, v := range service.services {
		addrs = append(addrs, v)
	}
	return addrs
}

func (service *ServiceDiscovery) Close() error {
	return service.cli.Close()
}

func main() {
	endpoints := []string{"localhost:2379"}
	service := NewServiceDiscovery(endpoints)
	defer service.cli.Close()

	service.WatchServices("/web/")
	service.WatchServices("/grpc/")

	for range time.Tick(time.Second * 5) {
		log.Println(service.GetServices())
	}
}

package discovery

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
	"log"
	"sync"
	"time"
)

const schema = "ns"

type ServiceDiscovery struct {
	lock     sync.Mutex
	cli      *clientv3.Client
	rcc      resolver.ClientConn
	services map[string]resolver.Address
}

func New(endpoints []string) resolver.Builder {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceDiscovery{
		cli: cli,
	}
}

func (sd *ServiceDiscovery) Build(target resolver.Target, rcc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	sd.rcc = rcc
	sd.services = make(map[string]resolver.Address)
	prefix := fmt.Sprintf("/%s/%s/", target.Scheme, target.Endpoint)
	resp, err := sd.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, kv := range resp.Kvs {
		sd.Set(string(kv.Key), string(kv.Value))
	}
	go sd.watch(prefix)
	return sd, nil
}

// ResolveNow 监视目标更新
func (sd *ServiceDiscovery) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}

// Scheme return schema
func (sd *ServiceDiscovery) Scheme() string {
	return schema
}

// Close 关闭
func (sd *ServiceDiscovery) Close() {
	log.Println("Close")
	err := sd.cli.Close()
	if err != nil {
		return
	}
}

func (sd *ServiceDiscovery) Set(k, v string) {
	sd.lock.Lock()
	defer sd.lock.Unlock()
	sd.services[k] = resolver.Address{Addr: v}
	addresses := make([]resolver.Address, 0, len(sd.services))
	for _, v := range sd.services {
		addresses = append(addresses, v)
	}
	sd.rcc.UpdateState(resolver.State{
		Addresses: addresses,
	})
}

func (sd *ServiceDiscovery) watch(prefix string) {
	log.Printf("watching prefix: %s", prefix)
	watch := sd.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for resp := range watch {
		for _, ev := range resp.Events {
			switch ev.Type {
			case 0:
				sd.Set(string(ev.Kv.Key), string(ev.Kv.Value))
			case 1:
				log.Println("TODO")
			}
		}
	}
}

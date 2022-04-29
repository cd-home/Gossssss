package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type ServiceRegister struct {
	cli           *clientv3.Client
	leaseID       clientv3.LeaseID
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
	val           string
}

func NewServiceRegister(endpoints []string, key, val string, lease int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	service := &ServiceRegister{
		cli: cli,
		key: key,
		val: val,
	}
	// set lease
	resp, err := cli.Grant(context.Background(), lease)
	if err != nil {
		return nil, err
	}

	// register service and band lease
	_, err = cli.Put(context.Background(), service.key, service.val, clientv3.WithLease(resp.ID))
	if err != nil {
		return nil, err
	}

	leaseRespChan, err := cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return nil, err
	}
	service.leaseID = resp.ID
	service.keepAliveChan = leaseRespChan

	return service, nil
}

func (service *ServiceRegister) ListenLeaseRespChan() {
	for leaseRespChan := range service.keepAliveChan {
		log.Println("续租成功", leaseRespChan)
	}
	log.Println("关闭续租")
}

func (service *ServiceRegister) CancelService() error {
	if _, err := service.cli.Revoke(context.Background(), service.leaseID); err != nil {
		return err
	}
	log.Println("注销租约")
	return service.cli.Close()
}

func main() {
	endpoints := []string{"localhost:2379"}
	service, err := NewServiceRegister(endpoints, "/web/node1", "localhost:8080", 5)
	if err != nil {
		log.Fatalln(err)
	}
	service.ListenLeaseRespChan()
}

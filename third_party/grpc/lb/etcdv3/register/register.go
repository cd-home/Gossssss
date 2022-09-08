package register

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

const schema = "ns"

type ServiceRegister struct {
	cli         *clientv3.Client
	leaseID     clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
	k           string
	v           string
}

func New(endpoints []string, service, addr string, lease int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	sr := &ServiceRegister{
		cli: cli,
		k:   fmt.Sprintf("/%s/%s/%s", schema, service, addr),
		v:   addr,
	}
	err = sr.Put(lease)
	if err != nil {
		return nil, err
	}
	return sr, nil
}

func (sr *ServiceRegister) Put(lease int64) error {
	resp, err := sr.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	_, err = sr.cli.Put(context.Background(), sr.k, sr.v, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	alive, err := sr.cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return err
	}
	sr.leaseID = resp.ID
	sr.keepAliveCh = alive
	log.Printf("PUT key: %s  val:%s  success!", sr.k, sr.v)
	return nil
}

func (sr *ServiceRegister) Revoke() error {
	//撤销租约
	if _, err := sr.cli.Revoke(context.Background(), sr.leaseID); err != nil {
		return err
	}
	return sr.cli.Close()
}

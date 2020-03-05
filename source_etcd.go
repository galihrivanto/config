package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	etcd "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type sourceEtcd struct {
	host   string
	prefix string
	sync.RWMutex

	client *etcd.Client

	// decoder
	decoder Decoder

	// current changeset
	current *Snapshot
}

func (s *sourceEtcd) Load() (*Snapshot, error) {
	return s.readKey()

	// return &Snapshot{}, nil
}

func (s *sourceEtcd) SetDecoder(decoder Decoder) {
	s.decoder = decoder
}

func (s *sourceEtcd) Watch(ctx context.Context) {

}

func (s *sourceEtcd) readKey() (*Snapshot, error) {
	log.Println("read key", s.prefix)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	rsp, err := s.client.Get(ctx, s.prefix, etcd.WithPrefix())
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	log.Println("resp", rsp)

	if rsp == nil || len(rsp.Kvs) == 0 {
		return nil, fmt.Errorf("source not found: %s", s.prefix)
	}

	kvs := make([]*mvccpb.KeyValue, 0, len(rsp.Kvs))
	for _, v := range rsp.Kvs {
		kvs = append(kvs, (*mvccpb.KeyValue)(v))
	}

	fmt.Println("kvs", kvs)

	return &Snapshot{}, nil
}

// Etcd create etcd source loader
func Etcd(endpoints string, dialTimeout time.Duration, vars ...string) Loader {
	// create etcd client
	var username, password, prefix string

	if len(vars) > 0 {
		prefix = vars[0]
	}

	if len(vars) == 3 {
		username = vars[1]
		password = vars[2]
	}

	addrs := strings.Split(endpoints, ",")
	if len(addrs) == 0 {
		addrs = []string{"localhost:2379"}
	}

	clientConfig := etcd.Config{
		Endpoints:   addrs,
		DialTimeout: dialTimeout,
		Username:    username,
		Password:    password,
	}

	log.Println("connecting")
	c, err := etcd.New(clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done..")

	return &sourceEtcd{
		client: c,
		prefix: prefix,
	}
}

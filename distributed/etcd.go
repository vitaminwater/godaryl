package distributed

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/config"
)

var pool sync.Pool

func SetKey(key, value string) error {
	cli := pool.Get().(*clientv3.Client)

	_, err := cli.Put(context.Background(), key, value)
	if err != nil {
		cli.Close()
		return err
	}
	defer pool.Put(cli)
	return nil
}

func GetKey(key string) (string, error) {
	cli := pool.Get().(*clientv3.Client)

	resp, err := cli.Get(context.Background(), key)
	if err != nil {
		cli.Close()
		return "", err
	}
	defer pool.Put(cli)
	if len(resp.Kvs) == 0 {
		return "", errors.New("Key not found")
	}
	value := string(resp.Kvs[0].Value)
	return value, nil
}

func ListPrefix(prefix string) ([]string, error) {
	cli := pool.Get().(*clientv3.Client)

	resp, err := cli.Get(context.Background(), "private_", clientv3.WithPrefix())
	if err != nil {
		cli.Close()
		return nil, err
	}
	defer pool.Put(cli)

	values := []string{}
	for _, kv := range resp.Kvs {
		values = append(values, string(kv.Value))
	}
	return values, nil
}

func Beacon(key, value string) {
	for {
		cli := pool.Get().(*clientv3.Client)
		gr, err := cli.Grant(context.TODO(), 2)
		if err != nil {
			log.Warning("Beacon ", err)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		_, err = cli.Put(ctx, key, value, clientv3.WithLease(gr.ID))
		cancel()
		if err != nil {
			log.Warning("Beacon ", err)
			continue
		}
		pool.Put(cli)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func Init() {
	pool = sync.Pool{
		New: func() interface{} {
			c, err := clientv3.New(clientv3.Config{
				Endpoints:   []string{config.AppContext.String("etcd-url")},
				DialTimeout: 5 * time.Second,
			})
			if err != nil {
				log.Fatal("distributed.Init ", err)
			}
			return c
		},
	}
}

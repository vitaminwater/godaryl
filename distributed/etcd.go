package distributed

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/vitaminwater/daryl/config"
)

var pool sync.Pool

func Beacon(key, value string) {
	for {
		cli := pool.Get().(*clientv3.Client)
		defer pool.Put(cli)
		gr, err := cli.Grant(context.TODO(), 2)
		if err != nil {
			log.Fatal("Beacon ", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		_, err = cli.Put(ctx, key, value, clientv3.WithLease(gr.ID))
		cancel()
		if err != nil {
			log.Fatal("Beacon ", err)
		}
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

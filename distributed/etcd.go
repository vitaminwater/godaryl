package distributed

import (
	"context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/vitaminwater/daryl/config"
)

var cli *clientv3.Client

func Beacon(key, value string) {
	for {
		gr, err := cli.Grant(context.TODO(), 2)
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = cli.Put(ctx, key, value, clientv3.WithLease(gr.ID))
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func Init() {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.AppContext.String("etcd-url")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	cli = c
}

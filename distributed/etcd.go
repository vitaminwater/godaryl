package distributed

import (
	"context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func Beacon(key, value string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	gr, err := cli.Grant(context.TODO(), 2)
	if err != nil {
		log.Fatal(err)
	}

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = cli.Put(ctx, key, value, clientv3.WithLease(gr.ID))
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

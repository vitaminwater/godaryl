package distributed

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/vitaminwater/daryl/config"
)

func FindDarylServer() func(string) (string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.AppContext.String("etcd-url")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	return func(identifier string) (string, error) {
		resp, err := cli.Get(context.Background(), fmt.Sprintf("daryl_%s", identifier))
		if err != nil {
			return "", err
		}
		if len(resp.Kvs) != 0 {
			url := string(resp.Kvs[0].Value)
			return url, nil
		}
		return "", errors.New("Daryl not found")
	}
}

func ListDarylServers() func() ([]string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.AppContext.String("etcd-url")},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	return func() ([]string, error) {
		resp, err := cli.Get(context.Background(), "private_", clientv3.WithPrefix())
		if err != nil {
			return nil, err
		}
		if len(resp.Kvs) != 0 {
			urls := []string{}
			for _, kv := range resp.Kvs {
				urls = append(urls, string(kv.Value))
			}
			return urls, nil
		}
		return nil, errors.New("No daryl servers")
	}
}

package config

import "github.com/urfave/cli"

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "advertized-url, u",
		Value: "localhost:8081",
		Usage: "Advertized URL",
	},
	cli.StringFlag{
		Name:  "bind-string, b",
		Value: "localhost:8081",
		Usage: "Advertized IP",
	},
	cli.StringFlag{
		Name:  "etcd-url, e",
		Value: "localhost:2379",
		Usage: "Advertized IP",
	},
	cli.StringFlag{
		Name:  "postgres-url, p",
		Value: "host=localhost user=daryl password=daryl dbname=daryl sslmode=disable",
		Usage: "Advertized IP",
	},
	cli.StringFlag{
		Name:  "redis-url, r",
		Value: "localhost:6379",
		Usage: "Advertized IP",
	},
}

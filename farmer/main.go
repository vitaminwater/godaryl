package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/config"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/distributed"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

func startDaryl(servers []string, d model.Daryl) {
	s := servers[int(float64(len(servers))*rand.Float64())]
	c, cl := protodef.OpenFarmConnection(s)
	defer cl()
	dp, err := d.ToProtodef()
	if err != nil {
		log.Fatal(err)
	}
	c.StartDaryl(context.Background(), &protodef.StartDarylRequest{Daryl: dp})
}

func startDaryls(c *cli.Context) {
	log.Info("startDaryl")
	for {
		daryls := []model.Daryl{}
		if err := daryl_db.Select("daryl", "", &daryls, model.Daryl{}); err != nil {
			log.Fatal(err)
		}

		servers, err := distributed.ListDarylServers()
		if err != nil {
			if err.Error() == "No daryl servers" {
				time.Sleep(10 * time.Second)
				continue
			} else {
				log.Fatal(err)
			}
		}
		for _, d := range daryls {
			_, err := distributed.FindDarylServer(d.Id)
			if err != nil {
				if err.Error() == "Key not found" {
					startDaryl(servers, d)
				} else {
					log.Fatal(err)
				}
				continue
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Daryl"
	app.Usage = "Show me what you got"
	app.Flags = []cli.Flag{
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
	app.Action = func(c *cli.Context) error {
		config.AppContext = c
		daryl_db.Init(false)
		distributed.Init()
		startDaryls(c)
		return nil
	}
	app.Run(os.Args)
}

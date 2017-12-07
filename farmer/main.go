package main

import (
	"os"

	"github.com/labstack/gommon/log"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/protodef"
)

func startDaryl(client protodef.FarmServiceClient, c *cli.Context) {
	log.Info("startDaryl")
}

func main() {
	app := cli.NewApp()
	app.Name = "Daryl"
	app.Usage = "Show me what you got"

	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start a new daryl",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "identifier, i",
					Usage: "Daryl's identifier",
				},
				cli.StringFlag{
					Name:  "name, n",
					Value: "Daryl",
					Usage: "Daryl's name",
				},
				cli.StringFlag{
					Name:  "password, p",
					Usage: "Daryl's password",
				},
			},
			Action: func(c *cli.Context) error {
				farm, _ := openConnection(c)
				startDaryl(farm, c)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

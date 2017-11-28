package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/config"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/distributed"
	"github.com/vitaminwater/daryl/kv"
	"github.com/vitaminwater/daryl/protodef"
)

func startServer() {
	router := gin.Default()
	pr := router.Group("/public")
	{
		pr.POST("/daryl", handleCreateDaryl)
	}
	dr := router.Group("/daryl")
	{
		dr.Use(setDarylServer())
		dr.POST("/cmd/:command", handleHTTPCommand)
	}
	router.Run()
}

func handleCreateDaryl(c *gin.Context) {
	d := &protodef.Daryl{}
	if err := c.Bind(d); err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}
	err := daryl_db.Insert("daryl", d)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}
	t, err := newTokenForDaryl(d)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}
	f := openFarmConnection("localhost:8081")
	f.StartDaryl(context.Background(), &protodef.StartDarylRequest{DarylIdentifier: d.Id})
	c.JSON(200, gin.H{
		"status": "ok",
		"daryl":  gin.H{"id": d.Id},
		"token":  gin.H{"hash": t.Hash},
	})
}

func main() {
	app := cli.NewApp()
	app.Name = "Daryl public server"
	app.Usage = "Show me what you got"
	app.Flags = config.Flags
	app.Action = func(c *cli.Context) error {
		config.AppContext = c
		distributed.Init()
		daryl_db.Init()
		kv.Init()
		startServer()
		return nil
	}
	app.Run(os.Args)
}

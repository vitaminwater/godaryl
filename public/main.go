package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/config"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/distributed"
	"github.com/vitaminwater/daryl/protodef"
)

const AUTH_TOKEN_HEADER = "X-Daryl-Auth-Token"

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
		return
	}
	err := daryl_db.Insert("daryl", d)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		return
	}
	c.JSON(200, gin.H{})
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
		startServer()
		return nil
	}
	app.Run(os.Args)
}

package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"github.com/vitaminwater/daryl/config"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/distributed"
	"github.com/vitaminwater/daryl/kv"
)

func startServer() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "POST", "OPTION"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	pr := router.Group("/public")
	{
		pr.POST("/daryl/token", handleCreateDarylToken)
		pr.POST("/daryl", handleCreateDaryl)
	}
	dr := router.Group("/daryl")
	{
		dr.Use(setDarylServer())
		dr.POST("/cmd/:command", handleHTTPCommand)
		dr.GET("/cmd/:command", handleHTTPCommand)
		dr.GET("/stream/:token", handleWS)
	}
	router.Run(config.AppContext.String("bind-string"))
}

func main() {
	app := cli.NewApp()
	app.Name = "Daryl public server"
	app.Usage = "Show me what you got"
	app.Flags = config.Flags
	app.Action = func(c *cli.Context) error {
		config.AppContext = c
		distributed.Init()
		kv.Init()
		daryl_db.Init(false)
		startServer()
		return nil
	}
	app.Run(os.Args)
}

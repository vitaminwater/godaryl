package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func processWS(conn *websocket.Conn, c *gin.Context) {
	id := c.MustGet("daryl_id")
	conn.WriteJSON(gin.H{"LOL": "LOLOL", "daryl": id})
	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			log.Info(err)
			return
		}
		log.Info(t, string(p))
	}
}

func handleWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Info("WS CONNECTED")
	go processWS(conn, c.Copy())
}

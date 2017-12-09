package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/kv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func subscribeRedis(conn *websocket.Conn, c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Warn(err)
		}
	}()

	id := c.MustGet("daryl_id")
	k := kv.Pool.Get()
	pkv := redis.PubSubConn{Conn: k}
	pkv.PSubscribe(fmt.Sprintf("daryl.%s.*", id))
	defer pkv.Close()
	for {
		switch v := pkv.Receive().(type) {
		case redis.Message:
			var m interface{}
			err := json.Unmarshal(v.Data, &m)
			if err != nil {
				log.Info(err)
			} else {
				conn.WriteJSON(m)
			}
		case redis.PMessage:
			var m interface{}
			err := json.Unmarshal(v.Data, &m)
			if err != nil {
				log.Info(err)
			} else {
				conn.WriteJSON(m)
			}
		case redis.Subscription:
		case error:
			log.Info(v)
		}
	}
}

func processWS(conn *websocket.Conn, c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Warn(err)
		}
	}()

	id := c.MustGet("daryl_id")
	conn.WriteJSON(gin.H{"daryl": id})
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
		log.Info(err)
		return
	}
	log.Info("WS CONNECTED")
	go processWS(conn, c.Copy())
	go subscribeRedis(conn, c.Copy())
}

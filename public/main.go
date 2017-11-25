package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/vitaminwater/daryl/protodef"

	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const AUTH_TOKEN_HEADER = "X-Daryl-Auth-Token"

func userMessage(c *gin.Context) {
	url := c.MustGet("daryl_url").(string)
	m := protodef.Message{}
	if err := c.BindJSON(&m); err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	d := openDarylConnection(url)
	um := protodef.UserMessageRequest{Identifier: "lol", Message: &m}
	resp, err := d.UserMessage(context.Background(), &um)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Info(url, m, resp)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func addHabit(c *gin.Context) {
	url := c.MustGet("daryl_url")
	h := protodef.Habit{}
	if err := c.BindJSON(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Info(url, h)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func startWorkSession(c *gin.Context) {
	url := c.MustGet("daryl_url")
	s := protodef.StartWorkSessionRequest{}
	if err := c.BindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	log.Info(url, s)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func cancelWorkSession(c *gin.Context) {
	url := c.MustGet("daryl_url")
	log.Info(url)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func refuseWorkSession(c *gin.Context) {
	url := c.MustGet("daryl_url")
	log.Info(url)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}

func findDarylServer(identifier string) (string, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	resp, err := cli.Get(context.Background(), fmt.Sprintf("daryl_%s", identifier))
	if err != nil {
		log.Fatal(err)
	}
	if len(resp.Kvs) != 0 {
		url := string(resp.Kvs[0].Value)
		log.Info(url)
		return url, nil
	}
	return "", errors.New("Daryl not found")
}

func setDarylServer(c *gin.Context) {
	url, err := findDarylServer("lol")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
		return
	}
	log.Infof("Daryl at %s", url)
	c.Set("daryl_url", url)
	c.Next()
}

func main() {
	router := gin.Default()
	router.Use(setDarylServer)
	router.POST("/message", userMessage)
	router.POST("/habit", addHabit)

	session := router.Group("/session")
	{
		session.POST("/", startWorkSession)
		session.POST("/cancel", cancelWorkSession)
		session.GET("/refuse", refuseWorkSession)
	}
	router.Run()
}

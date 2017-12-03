package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/config"
)

func findDarylServer() func(string) (string, error) {
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

func setDarylServer() func(*gin.Context) {
	fds := findDarylServer()
	return func(c *gin.Context) {
		h := c.GetHeader(AUTH_TOKEN_HEADER)
		if h == "" {
			h = c.Param("token")
		}
		if h == "" {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": errors.New("Access denied")})
			c.Abort()
			return
		}
		t, err := newTokenFromToken(h)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": err})
			c.Abort()
			return
		}
		url, err := fds(t.Daryl.Id)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": err})
			c.Abort()
			return
		}
		log.Infof("Daryl at %s", url)
		c.Set("daryl_url", url)
		c.Set("daryl_id", t.Daryl.Id)
	}
}

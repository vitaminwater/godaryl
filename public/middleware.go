package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/distributed"
)

func setDarylServer() func(*gin.Context) {
	findDarylServer := distributed.FindDarylServer()
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
		url, err := findDarylServer(t.Daryl.Id)
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

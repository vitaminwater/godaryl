package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitaminwater/daryl/distributed"
)

func setDarylServer() func(*gin.Context) {
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
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": err.Error()})
			c.Abort()
			return
		}
		url, err := distributed.FindDarylServer(t.Daryl.Id)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "error": err.Error()})
			c.Abort()
			return
		}
		c.Set("daryl_url", url)
		c.Set("daryl_id", t.Daryl.Id)
	}
}

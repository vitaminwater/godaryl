package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/vitaminwater/daryl/distributed"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

const AUTH_TOKEN_HEADER = "X-Daryl-Auth-Token"

type token struct {
	Hash  string          `json:"hash"`
	Daryl *protodef.Daryl `json:"daryl"`
}

func (t *token) save() error {
	err := distributed.SetKey(t.Hash, t.Daryl.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *token) load() error {
	id, err := distributed.GetKey(t.Hash)
	if err != nil {
		return err
	}

	t.Daryl.Id = id
	return nil
}

func newTokenForDaryl(d *protodef.Daryl) (*token, error) {
	u1 := uuid.NewV4()
	t := &token{
		Hash:  fmt.Sprintf("%s", u1),
		Daryl: d,
	}
	t.save()
	return t, nil
}

func newTokenFromToken(hash string) (*token, error) {
	t := &token{Hash: hash, Daryl: &protodef.Daryl{}}
	err := t.load()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func handleCreateDarylToken(c *gin.Context) {
	d := &protodef.Daryl{}
	if err := c.Bind(d); err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	da, err := model.NewDarylFromProtodef(d)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	if err := da.GetFromEmailAndPassword(); err != nil {
		c.JSON(500, gin.H{"status": "error", "error": err})
		c.Abort()
		return
	}

	d, err = da.ToProtodef()
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
	c.JSON(200, gin.H{
		"status": "ok",
		"token":  gin.H{"hash": t.Hash, "daryl": d},
	})
}

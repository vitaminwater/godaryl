package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/vitaminwater/daryl/kv"
	"github.com/vitaminwater/daryl/protodef"
)

const AUTH_TOKEN_HEADER = "X-Daryl-Auth-Token"

type token struct {
	Hash  string          `json:"hash"`
	Daryl *protodef.Daryl `json:"daryl"`
}

func (t *token) save() error {
	conn := kv.Pool.Get()
	defer conn.Close()

	conn.Send("HSET", t.Hash, "hash", t.Hash)
	conn.Send("HSET", t.Hash, "daryl", t.Daryl.Id)
	conn.Flush()
	return nil
}

func (t *token) load() error {
	conn := kv.Pool.Get()
	defer conn.Close()

	resp, err := redis.String(conn.Do("HGET", t.Hash, "daryl"))
	if err != nil {
		return err
	}

	t.Daryl.Id = resp
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

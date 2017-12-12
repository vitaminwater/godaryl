package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/vitaminwater/daryl/distributed"
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

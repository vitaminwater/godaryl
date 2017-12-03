package model

import (
	"encoding/json"

	"github.com/jmoiron/sqlx/types"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/protodef"
)

type Message struct {
	Id      string `json:"id" db:"id" access:"s"`
	Text    string `json:"text" db:"text" access:"i,u,s"`
	DarylId string `json:"darylId" db:"daryl_id" access:"i,s"`
	HabitId string `json:"habit_id" db:"habit_id" access:"i,s"`

	AttrsDB types.JSONText         `json:"-" db:"attrs" access:"i,u,s"`
	Attrs   map[string]interface{} `json:"attrs" db:"-"`
}

func (m *Message) Insert() error {
	a, err := json.Marshal(m.Attrs)
	if err != nil {
		return err
	}
	m.AttrsDB = a
	return daryl_db.Insert("message", m)
}

func (m Message) Update() error {
	a, err := json.Marshal(m.Attrs)
	if err != nil {
		return err
	}
	m.AttrsDB = a
	return daryl_db.Update("message", "id", m)
}

func (m Message) ToProtodef() (*protodef.Message, error) {
	a, err := json.Marshal(m.Attrs)
	if err != nil {
		return nil, err
	}
	return &protodef.Message{
		Id:    m.Id,
		Text:  m.Text,
		Attrs: a,
	}, nil
}

func NewMessageFromProtodef(d Daryl, msg *protodef.Message) Message {
	m := Message{
		Id:      msg.Id,
		Text:    msg.Text,
		DarylId: d.Id,
		HabitId: msg.HabitIdentifier,
		Attrs:   make(map[string]interface{}),
	}
	return m
}

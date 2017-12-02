package model

import (
	"github.com/vitaminwater/daryl/protodef"
)

type Message struct {
	Id      string `json:"id,omitempty" db:"id" access:"s"`
	Text    string `json:"text,omitempty" db:"text" access:"i,u,s"`
	DarylId string `json:"darylId" db:"daryl_id" access:"i,s"`
	HabitId string `json:"habit_id" db:"habit_id" access:"i,s"`

	Attrs map[string]interface{} `json:"attrs" db:"attrs" access:"i,u,s"`
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

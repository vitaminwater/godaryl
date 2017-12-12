package model

import (
	"encoding/json"

	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/protodef"
)

type Message struct {
	Id      string `json:"id" db:"id" access:"s"`
	Text    string `json:"text" db:"text" access:"i,u,s"`
	DarylId string `json:"darylId" db:"daryl_id" access:"i,s"`
	HabitId string `json:"habitId" db:"habit_id" access:"i,s"`

	Attrs daryl_db.PropertyMap `json:"attrs" db:"attrs" access:"i,u,s"`
}

func (m *Message) Insert() error {
	return daryl_db.Insert("message", m)
}

func (m Message) Update() error {
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

func NewMessageFromProtodef(d Daryl, msg *protodef.Message) (Message, error) {
	attrs := map[string]interface{}{}
	err := json.Unmarshal(msg.Attrs, &attrs)
	if err != nil {
		return Message{}, err
	}

	m := Message{
		Id:      msg.Id,
		Text:    msg.Text,
		DarylId: d.Id,
		HabitId: msg.HabitIdentifier,
		Attrs:   attrs,
	}
	return m, nil
}

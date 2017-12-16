package model

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/protodef"
)

type Message struct {
	Id      string         `json:"id" db:"id" access:"s"`
	Text    string         `json:"text" db:"text" access:"i,u,s"`
	At      time.Time      `json:"at" db:"at"`
	DarylId string         `json:"darylId" db:"daryl_id" access:"i,s"`
	HabitId sql.NullString `json:"habitId" db:"habit_id" access:"i,s"`

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

	at, err := ptypes.TimestampProto(m.At)
	if err != nil {
		return nil, err
	}
	return &protodef.Message{
		Id:    m.Id,
		Text:  m.Text,
		At:    at,
		Attrs: a,
	}, nil
}

func NewMessageFromProtodef(d Daryl, msg *protodef.Message) (Message, error) {
	attrs := map[string]interface{}{}
	err := json.Unmarshal(msg.Attrs, &attrs)
	if err != nil {
		return Message{}, err
	}

	var habitId sql.NullString
	if msg.HabitIdentifier != "" {
		habitId.Scan(msg.HabitIdentifier)
	}

	at, err := ptypes.Timestamp(msg.At)
	if err != nil {
		return Message{}, err
	}

	m := Message{
		Id:      msg.Id,
		Text:    msg.Text,
		At:      at,
		DarylId: d.Id,
		HabitId: habitId,
		Attrs:   attrs,
	}
	return m, nil
}

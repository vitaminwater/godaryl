package model

import (
	"encoding/json"

	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/protodef"
)

type Trigger struct {
	Id      string               `json:"id" db:"id" access:"s"`
	HabitId string               `json:"habitId" db:"habit_id" access:"i,s"`
	DarylId string               `json:"darylId" db:"daryl_id" access:"i,s"`
	Name    string               `json:"name" db:"name" access:"i,s"`
	Engine  string               `json:"engine" db:"engine" access:"i,s"`
	Params  daryl_db.PropertyMap `json:"params" db:"params" access:"i,u,s"`
}

func (t *Trigger) Insert() error {
	return daryl_db.Insert("habit_trigger", t)
}

func (t Trigger) Update() error {
	return daryl_db.Update("habit_trigger", "id", t)
}

func (t Trigger) ToProtodef() (*protodef.Trigger, error) {
	p, err := json.Marshal(t.Params)
	if err != nil {
		return nil, err
	}

	return &protodef.Trigger{
		Id:              t.Id,
		HabitIdentifier: t.HabitId,
		Name:            t.Name,
		Engine:          t.Engine,
		Params:          p,
	}, nil
}

func TriggersForDaryl(d Daryl) ([]Trigger, error) {
	result := []Trigger{}
	err := daryl_db.Select("habit_trigger", "daryl_id", &result, Trigger{DarylId: d.Id})
	if err != nil {
		return result, err
	}

	return result, nil
}

func NewTriggerFromProtodef(h Habit, t *protodef.Trigger) (Trigger, error) {
	params := map[string]interface{}{}
	err := json.Unmarshal(t.Params, &params)
	if err != nil {
		return Trigger{}, err
	}

	return Trigger{
		Id:      t.Id,
		HabitId: t.HabitIdentifier,
		DarylId: h.DarylId,
		Name:    t.Name,
		Engine:  t.Engine,
		Params:  params,
	}, nil
}

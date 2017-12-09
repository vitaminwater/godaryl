package model

import (
	"encoding/json"

	"github.com/jmoiron/sqlx/types"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/protodef"
)

type Trigger struct {
	Id       string                 `json:"id" db:"id" access:"s"`
	HabitId  string                 `json:"habitId" db:"habit_id" access:"i,s"`
	Name     string                 `json:"name" db:"name" access:"i,s"`
	Engine   string                 `json:"engine" db:"engine" access:"i,s"`
	ParamsDB types.JSONText         `json:"-" db:"params" access:"i,u,s"`
	Params   map[string]interface{} `json:"params" db:"-"`
}

func (t *Trigger) Insert() error {
	a, err := json.Marshal(t.Params)
	if err != nil {
		return err
	}
	t.ParamsDB = a

	return daryl_db.Insert("habit_trigger", t)
}

func (t Trigger) Update() error {
	a, err := json.Marshal(t.Params)
	if err != nil {
		return err
	}
	t.ParamsDB = a

	return nil
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

func NewTriggerFromProtodef(h Habit, t *protodef.Trigger) (Trigger, error) {
	params := map[string]interface{}{}
	err := json.Unmarshal(t.Params, &params)
	if err != nil {
		return Trigger{}, err
	}

	return Trigger{
		Id:      t.Id,
		HabitId: t.HabitIdentifier,
		Name:    t.Name,
		Engine:  t.Engine,
		Params:  params,
	}, nil
}

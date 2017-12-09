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
	return daryl_db.Insert("habit_trigger", t)
}

func (h Trigger) Update() error {
	return nil
}

func (t Trigger) ToProtodef() (*protodef.Trigger, error) {
	p, err := json.Marshal(t.Params)
	if err != nil {
		return nil, err
	}

	return &protodef.Trigger{
		Id:     t.Id,
		Name:   t.Name,
		Engine: t.Engine,
		Params: p,
	}, nil
}

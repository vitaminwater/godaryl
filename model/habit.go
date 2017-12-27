package model

import (
	"time"

	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/protodef"
)

type Habit struct {
	Id       string        `json:"id" db:"id" access:"s"`
	Title    string        `json:"title" db:"title" access:"i,u,s"`
	Duration time.Duration `json:"duration" db:"duration"`
	DarylId  string        `json:"darylId" db:"daryl_id" access:"i,s"`
}

func (h *Habit) Insert() error {
	return daryl_db.Insert("habit", h)
}

func (h Habit) Update() error {
	return nil
}

func (h Habit) ToProtodef() (*protodef.Habit, error) {
	return &protodef.Habit{
		Id:       h.Id,
		Title:    h.Title,
		Duration: h.Duration.String(),
	}, nil
}

func HabitsForDaryl(darylId string) ([]Habit, error) {
	result := []Habit{}
	err := daryl_db.Select("habit", "daryl_id", &result, Habit{DarylId: darylId})
	if err != nil {
		return result, err
	}

	return result, nil
}

func NewHabitFromProtodef(darylId string, h *protodef.Habit) (Habit, error) {
	duration, err := time.ParseDuration(h.Duration)
	if err != nil {
		return Habit{}, err
	}
	return Habit{
		Id:       h.Id,
		Title:    h.Title,
		DarylId:  darylId,
		Duration: duration,
	}, nil
}

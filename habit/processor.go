package habit

import (
	"github.com/vitaminwater/daryl/daryl"
)

type habitProcessor struct {
	d     *daryl.Daryl
	store *habitStore
}

func (hp *habitProcessor) SetDaryl(d *daryl.Daryl) {
	hp.d = d
	hp.store = newHabitStore(d)
}

func (hp *habitProcessor) AddHabit(r *daryl.AddHabitRequest) (*daryl.AddHabitResponse, error) {
	hp.d.Pub(r, daryl.ADD_HABIT_TOPIC)
	return &daryl.AddHabitResponse{}, nil
}

func NewHabitProcessor() *habitProcessor {
	hp := &habitProcessor{}
	return hp
}

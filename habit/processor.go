package habit

import (
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type habitProcessor struct {
	d     *daryl.Daryl
	store *habitStore
}

func (hp *habitProcessor) SetDaryl(d *daryl.Daryl) {
	hp.d = d
	hp.store = newHabitStore(d)
}

func (hp *habitProcessor) AddHabit(r *protodef.AddHabitRequest) (*protodef.AddHabitResponse, error) {
	hp.d.Pub(r, daryl.ADD_HABIT_TOPIC)
	return &protodef.AddHabitResponse{}, nil
}

func NewHabitProcessor() *habitProcessor {
	hp := &habitProcessor{}
	return hp
}

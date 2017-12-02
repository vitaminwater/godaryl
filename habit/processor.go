package habit

import (
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
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
	h, err := model.NewHabitFromProtodef(hp.d.D, r.Habit)
	if err != nil {
		return nil, err
	}
	hp.store.addHabit(h)
	hp.d.Pub(h, daryl.ADD_HABIT_TOPIC)
	return &protodef.AddHabitResponse{Habit: r.Habit}, nil
}

func (hp *habitProcessor) GetDueHabits() []model.Habit {
	return hp.store.getDueHabits()
}

func (hp *habitProcessor) GetWeight(h model.Habit) int {
	return hp.store.getAttributes(h).Urgent
}

func NewHabitProcessor() *habitProcessor {
	hp := &habitProcessor{}
	return hp
}

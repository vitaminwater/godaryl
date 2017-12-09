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
	h = hp.store.addHabit(h)
	hp.d.Pub(h, daryl.ADD_HABIT_TOPIC)
	hh, err := h.ToProtodef()
	if err != nil {
		return nil, err
	}
	return &protodef.AddHabitResponse{Habit: hh}, nil
}

func (hp *habitProcessor) AddTrigger(r *protodef.AddTriggerRequest) (*protodef.AddTriggerResponse, error) {
	h, err := hp.store.getHabit(r.Trigger.HabitIdentifier)
	if err != nil {
		return nil, err
	}

	t, err := model.NewTriggerFromProtodef(h, r.Trigger)
	if err != nil {
		return nil, err
	}

	tp, err := t.ToProtodef()
	if err != nil {
		return nil, err
	}

	return &protodef.AddTriggerResponse{Trigger: tp}, nil
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

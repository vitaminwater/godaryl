package habit

import (
	log "github.com/sirupsen/logrus"
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
	h, err := model.NewHabitFromProtodef(hp.d.D.Id, r.Habit)
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

func (hp *habitProcessor) GetHabits(*protodef.GetHabitsRequest) (*protodef.GetHabitsResponse, error) {
	habits, err := hp.GetAllHabits()
	if err != nil {
		return nil, err
	}
	var hps []*protodef.Habit
	for _, h := range habits {
		hpp, err := h.GetHabit().ToProtodef()
		if err != nil {
			log.Warning(err)
			continue
		}
		hps = append(hps, hpp)
	}
	return &protodef.GetHabitsResponse{Habits: hps}, nil
}

func (hp *habitProcessor) GetHabit(id string) (daryl.Habit, error) {
	return hp.store.getHabit(id)
}

func (hp *habitProcessor) GetAllHabits() ([]daryl.Habit, error) {
	return hp.store.getHabits()
}

func (hp *habitProcessor) GetDueHabits() []daryl.Habit {
	return hp.store.getDueHabits()
}

func NewHabitProcessor() *habitProcessor {
	hp := &habitProcessor{}
	return hp
}

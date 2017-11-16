package habit

import (
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type habitProcessor struct {
	d     *daryl.Daryl
	store *habitStore
}

func habitArray(hs []*habit) []daryl.Habit {
	r := make([]daryl.Habit, 0)
	for _, h := range hs {
		r = append(r, h)
	}
	return r
}

func (hp *habitProcessor) SetDaryl(d *daryl.Daryl) {
	hp.d = d
	hp.store = newHabitStore(d)
}

func (hp *habitProcessor) AddHabit(r *protodef.AddHabitRequest) (*protodef.AddHabitResponse, error) {
	h := newHabit(r.Habit)
	hp.store.addHabit(h)
	hp.d.Pub(h, daryl.ADD_HABIT_TOPIC)
	return &protodef.AddHabitResponse{r.Habit}, nil
}

func (hp *habitProcessor) GetDueHabits() []daryl.Habit {
	r := make(chan []*habit)
	hp.store.c <- &storeCommandGetDueHabit{r}
	hs := <-r
	close(r)
	return habitArray(hs)
}

func NewHabitProcessor() *habitProcessor {
	hp := &habitProcessor{}
	return hp
}

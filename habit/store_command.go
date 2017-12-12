package habit

import (
	"github.com/labstack/gommon/log"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type storeCommand interface {
	execute(hs *habitStore)
}

type storeCommandAddHabit struct {
	h model.Habit
	r chan model.Habit
}

func (c *storeCommandAddHabit) execute(hs *habitStore) {
	if c.h.Id == "" {
		err := c.h.Insert()
		if err != nil {
			log.Info(err)
			return
		}
	}
	hw := newHabitWorker(hs.d, c.h)
	hs.habitWorkers.Store(c.h.Id, hw)
	c.r <- c.h
}

type storeCommandGetDueHabits struct {
	r chan []daryl.Habit
}

func (d *storeCommandGetDueHabits) execute(hs *habitStore) {
	habits := make([]daryl.Habit, 0, 10)
	hs.habitWorkers.Range(func(k, w interface{}) bool {
		a := w.(*habitWorker).getAttributes()
		if a.NMissed > 0 {
			habits = append(habits, w.(*habitWorker))
		}
		return true
	})
	d.r <- habits
}

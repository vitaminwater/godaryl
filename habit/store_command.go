package habit

import (
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/vitaminwater/daryl/model"
)

type storeCommand interface {
	execute(hs *habitStore)
}

type storeCommandGetAttrs struct {
	h model.Habit
	r chan Attributes
}

func (c *storeCommandGetAttrs) execute(hs *habitStore) {
	for _, w := range hs.habitWorkers {
		h := w.GetHabit()
		if h.h.Id == c.h.Id {
			c.r <- h.a
			return
		}
	}
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
	hs.habitWorkers[c.h.Id] = hw
	c.r <- c.h
}

type storeCommandGetDueHabits struct {
	r chan []model.Habit
}

func (d *storeCommandGetDueHabits) execute(hs *habitStore) {
	habits := make([]model.Habit, 0, 10)
	for _, w := range hs.habitWorkers {
		h := w.GetHabit()
		if h.a.NMissed > 0 {
			habits = append(habits, h.h)
		}
	}
	d.r <- habits
}

type storeCommandGetHabitResponse struct {
	h   model.Habit
	err error
}

type storeCommandGetHabit struct {
	id string
	r  chan storeCommandGetHabitResponse
}

func (c *storeCommandGetHabit) execute(hs *habitStore) {
	hw, ok := hs.habitWorkers[c.id]
	if ok == false {
		c.r <- storeCommandGetHabitResponse{err: errors.New("Habit not found")}
		return
	}
	h := hw.GetHabit()
	c.r <- storeCommandGetHabitResponse{h: h.h, err: nil}
}

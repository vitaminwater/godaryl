package habit

import (
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/vitaminwater/daryl/daryl"
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
	w := hs.habitWorkers[c.h.Id]
	c.r <- w.getAttributes()
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
	r chan []daryl.Habit
}

func (d *storeCommandGetDueHabits) execute(hs *habitStore) {
	habits := make([]daryl.Habit, 0, 10)
	for _, w := range hs.habitWorkers {
		a := w.getAttributes()
		if a.NMissed > 0 {
			habits = append(habits, w)
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
	c.r <- storeCommandGetHabitResponse{h: h, err: nil}
}

type storeCommandGetHabitWorkerResponse struct {
	hw  *habitWorker
	err error
}

type storeCommandGetHabitWorker struct {
	id string
	r  chan storeCommandGetHabitWorkerResponse
}

func (c *storeCommandGetHabitWorker) execute(hs *habitStore) {
	hw, ok := hs.habitWorkers[c.id]
	if ok == false {
		c.r <- storeCommandGetHabitWorkerResponse{err: errors.New("Habit not found")}
		return
	}
	c.r <- storeCommandGetHabitWorkerResponse{hw: hw, err: nil}
}

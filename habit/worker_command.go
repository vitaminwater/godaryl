package habit

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/vitaminwater/daryl/model"
)

type workerCommand interface {
	execute(worker *habitWorker)
}

type workerCommandOnHabitTrigger struct{}

func (oht *workerCommandOnHabitTrigger) execute(w *habitWorker) {
	w.a.NMissed++
	w.a.Urgent *= 2
	w.d.Pub(w.h, HABIT_SCHEDULED_TOPIC)
}

type workerCommandGetHabitResponse struct {
	h model.Habit
	a Attributes
}

type workerCommandGetHabit struct {
	r chan workerCommandGetHabitResponse
}

func (gh *workerCommandGetHabit) execute(w *habitWorker) {
	gh.r <- workerCommandGetHabitResponse{
		h: w.h,
		a: w.a,
	}
}

type workerCommandTick struct {
}

func (c *workerCommandTick) execute(w *habitWorker) {
	log.Info("tick")
	if w.a.NMissed > 0 {
		p := float64(time.Since(w.a.LastDone)/time.Minute) * float64(w.a.NMissed)
		w.a.Urgent += int(p)
	}
}

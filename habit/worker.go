package habit

import (
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
	"time"
)

type habit struct {
	protodef.Habit
	LastDone *time.Time
}

func newHabit(h *protodef.Habit) *habit {
	lastDone, err := ptypes.Timestamp(h.LastDone)
	if err != nil {
		log.Info(err)
	}
	return &habit{
		*h,
		&lastDone,
	}
}

type habitWorker struct {
	h *habit

	c chan interface{}
	d *daryl.Daryl
}

func habitWorkerProcess(h *habitWorker) {
	for msg := range h.c {
		log.Info("HABIT ", msg)
	}
}

func newHabitWorker(d *daryl.Daryl, h *protodef.Habit) *habitWorker {
	hw := &habitWorker{
		newHabit(h),
		d.Sub(
			daryl.ADD_HABIT_TOPIC,
		),
		d,
	}
	go habitWorkerProcess(hw)
	return hw
}

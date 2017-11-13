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
	Deadline *time.Time `db:"deadline"`
	LastDone *time.Time `db:"lastDone"`
}

func newHabit(h *protodef.Habit) *habit {
	lastDone, err := ptypes.Timestamp(h.LastDone)
	if err != nil {
		log.Info(err)
	}
	deadline, err := ptypes.Timestamp(h.Deadline)
	if err != nil {
		log.Info(err)
	}

	return &habit{
		*h,
		&deadline,
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

func newHabitWorker(d *daryl.Daryl, h *habit) *habitWorker {
	hw := &habitWorker{
		h,
		d.Sub(
			daryl.ADD_HABIT_TOPIC,
		),
		d,
	}
	go habitWorkerProcess(hw)
	return hw
}

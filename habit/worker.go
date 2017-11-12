package habit

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
	"time"
)

type habitFrequency struct {
	during    uint
	every     uint
	everyUnit string
}

type habit struct {
	title string

	avgDuration uint
	deadline    *time.Time
	frequency   habitFrequency

	lastDone *time.Time
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
	var deadline *time.Time = nil
	tDeadline := time.Now()
	if err := tDeadline.UnmarshalText([]byte(h.Deadline)); err != nil {
		deadline = &tDeadline
	}
	ha := &habit{
		h.Title,
		uint(h.AvgDuration),
		deadline,
		habitFrequency{
			uint(h.During),
			uint(h.Every),
			h.EveryUnit,
		},
		nil,
	}
	hw := &habitWorker{
		ha,
		d.Sub(
			daryl.ADD_HABIT_TOPIC,
		),
		d,
	}
	go habitWorkerProcess(hw)
	return hw
}

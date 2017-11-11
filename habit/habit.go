package habit

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
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

	c chan interface{}
}

func habitProcess(h *habit) {
	for msg := range h.c {
		log.Info("HABIT ", msg)
	}
}

func newHabit(h *daryl.Habit) *habit {
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
		make(chan interface{}, 10),
	}
	go habitProcess(ha)
	return ha
}

type habitWorker struct {
	d      *daryl.Daryl
	c      chan interface{}
	habits []*habit
}

func habitWorkerProcess(hw *habitWorker) {
	for msg := range hw.c {
		tm := msg.(daryl.TopicMessage)
		switch t := tm.Topic; t {
		case daryl.ADD_HABIT_TOPIC:
			r := tm.Msg.(*daryl.AddHabitRequest)
			h := newHabit(r.Habit)
			hw.habits = append(hw.habits, h)
		}
	}
}

func NewHabitWorker(d *daryl.Daryl) *habitWorker {
	hw := &habitWorker{d: d, habits: make([]*habit, 0, 10)}
	hw.c = d.Sub(
		daryl.ADD_HABIT_TOPIC,
	)
	go habitWorkerProcess(hw)
	return hw
}

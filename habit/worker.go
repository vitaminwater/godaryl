package habit

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
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

type workerCommand interface {
	execute(worker *habitWorker)
}

type workerCommandOnHabitTrigger struct{}

func (oht *workerCommandOnHabitTrigger) execute(w *habitWorker) {
	w.h.NMissed++
	log.Info("onHabitTrigger ", w.h.NMissed)
}

type habitWorker struct {
	h *habit

	cr   *cron.Cron
	tick <-chan time.Time
	cmd  chan workerCommand
	sub  chan interface{}
	d    *daryl.Daryl
}

func habitWorkerProcess(hw *habitWorker) {
	for {
		select {
		case t := <-hw.tick:
			log.Info("tick ", t)
		case cmd := <-hw.cmd:
			cmd.execute(hw)
		case msg := <-hw.sub:
			log.Info("sub ", msg)
		}
	}
}

func newHabitWorker(d *daryl.Daryl, h *habit) *habitWorker {
	hw := &habitWorker{
		h,
		cron.New(),
		time.Tick(time.Duration(1) * time.Minute),
		make(chan workerCommand, 10),
		d.Sub(
			daryl.ADD_HABIT_TOPIC,
		),
		d,
	}
	hw.cr.AddFunc(h.Cron, func() { hw.cmd <- &workerCommandOnHabitTrigger{} })
	go habitWorkerProcess(hw)
	hw.cr.Start()
	return hw
}

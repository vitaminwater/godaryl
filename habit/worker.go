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
	h        protodef.Habit
	Deadline *time.Time    `db:"deadline"`
	LastDone *time.Time    `db:"lastDone"`
	Duration time.Duration `db:"duration"`
}

func (h *habit) GetHabit() protodef.Habit {
	return h.h
}

func (h *habit) GetWeight() uint32 {
	return 0
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
	duration, err := time.ParseDuration(h.Duration)
	if err != nil {
		log.Info(err)
	}
	return &habit{
		*h,
		&deadline,
		&lastDone,
		duration,
	}
}

type workerCommand interface {
	execute(worker *habitWorker)
}

type workerCommandOnHabitTrigger struct{}

func (oht *workerCommandOnHabitTrigger) execute(w *habitWorker) {
	w.h.h.Stats.NMissed++
	w.h.h.Stats.Forget = uint64(float64(w.h.h.Stats.Forget) * 1.5)
	w.d.Pub(w.h, HABIT_SCHEDULED_TOPIC)
	log.Info("onHabitTrigger ", w.h.h.Stats.NMissed, " ", w.h.h.Stats.Forget)
}

type workerCommandGetHabit struct {
	r chan *habit
}

func (gh *workerCommandGetHabit) execute(w *habitWorker) {
	gh.r <- w.h
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
		case _ = <-hw.tick:
			p := float64(time.Since(*hw.h.LastDone)) / float64(hw.h.Duration) * 100
			hw.h.h.Stats.Forget += uint64(p)
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
	hw.cr.AddFunc(h.h.Cron, func() { hw.cmd <- &workerCommandOnHabitTrigger{} })
	go habitWorkerProcess(hw)
	hw.cr.Start()
	return hw
}

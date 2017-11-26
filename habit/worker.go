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
	Deadline *time.Time    `db:"deadline"`
	LastDone *time.Time    `db:"lastDone"`
	Duration time.Duration `db:"duration"`
}

func (h *habit) GetHabit() protodef.Habit {
	return h.Habit
}

func (h *habit) GetWeight() uint32 {
	return h.Stats.Urgent
}

func newHabit(h *protodef.Habit) *habit {
	log.Info(h)
	lastDone, err := ptypes.Timestamp(h.LastDone)
	if err != nil {
		log.Info(err)
		lastDone = time.Now()
	}
	deadline, err := ptypes.Timestamp(h.Deadline)
	if err != nil {
		log.Info(err)
	}
	duration, err := time.ParseDuration(h.Duration)
	if err != nil {
		log.Info(err)
	}
	h.Stats = &protodef.HabitStat{}
	return &habit{
		Habit:    *h,
		Deadline: &deadline,
		LastDone: &lastDone,
		Duration: duration,
	}
}

type workerCommand interface {
	execute(worker *habitWorker)
}

type workerCommandOnHabitTrigger struct{}

func (oht *workerCommandOnHabitTrigger) execute(w *habitWorker) {
	w.h.Stats.NMissed++
	w.h.Stats.Urgent *= 2
	w.d.Pub(w.h, HABIT_SCHEDULED_TOPIC)
}

type workerCommandGetHabit struct {
	r chan *habit
}

func (gh *workerCommandGetHabit) execute(w *habitWorker) {
	gh.r <- w.h
}

type worker interface {
	tick()
}

type habitWorker struct {
	h *habit

	cr  *cron.Cron
	t   <-chan time.Time
	cmd chan workerCommand
	sub chan interface{}
	d   *daryl.Daryl
}

func (hw *habitWorker) tick() {
	if hw.h.Stats.NMissed > 0 {
		p := float64(time.Since(*hw.h.LastDone)/time.Minute) * float64(hw.h.Stats.NMissed)
		hw.h.Stats.Urgent += uint32(p)
	}
}

func habitWorkerProcess(hw *habitWorker) {
	for {
		select {
		case _ = <-hw.t:
			hw.tick()
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
		time.Tick(time.Duration(10) * time.Minute),
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

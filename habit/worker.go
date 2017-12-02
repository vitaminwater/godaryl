package habit

import (
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type Attributes struct {
	Urgent   int        `json:"urgent"`
	NMissed  uint       `json:"n_missed"`
	LastDone *time.Time `json:"last_done"`
}

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

type habitWorker struct {
	d *daryl.Daryl
	a Attributes
	h model.Habit

	cr  *cron.Cron
	t   <-chan time.Time
	cmd chan workerCommand
	sub chan interface{}
}

func (hw *habitWorker) GetHabit() workerCommandGetHabitResponse {
	r := make(chan workerCommandGetHabitResponse)
	hw.cmd <- &workerCommandGetHabit{r}
	h := <-r
	close(r)
	return h
}

func (hw *habitWorker) tick() {
	if hw.a.NMissed > 0 {
		p := float64(time.Since(*hw.a.LastDone)/time.Minute) * float64(hw.a.NMissed)
		hw.a.Urgent += int(p)
	}
}

func habitWorkerProcess(hw *habitWorker) {
	log.Info("habitWorkerProcess")
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

func newHabitWorker(d *daryl.Daryl, h model.Habit) *habitWorker {
	hw := &habitWorker{
		d: d,
		h: h,

		cr:  cron.New(),
		t:   time.Tick(time.Duration(10) * time.Minute),
		cmd: make(chan workerCommand, 10),
		sub: d.Sub(
			daryl.ADD_HABIT_TOPIC,
		),
	}
	hw.cr.AddFunc(h.Cron, func() { hw.cmd <- &workerCommandOnHabitTrigger{} })
	go habitWorkerProcess(hw)
	hw.cr.Start()
	return hw
}

package habit

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type Attributes struct {
	Urgent   int       `json:"urgent"`
	NMissed  uint      `json:"n_missed"`
	LastDone time.Time `json:"last_done"`
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

type workerCommandTick struct {
}

func (c *workerCommandTick) execute(w *habitWorker) {
	log.Info("tick")
	if w.a.NMissed > 0 {
		p := float64(time.Since(w.a.LastDone)/time.Minute) * float64(w.a.NMissed)
		w.a.Urgent += int(p)
	}
}

type habitWorker struct {
	d *daryl.Daryl
	a Attributes
	h model.Habit

	cr  *cron.Cron
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

func habitWorkerProcess(hw *habitWorker) {
	defer func() {
		if err := recover(); err != nil {
			log.Warn(err)
			time.Sleep(time.Duration(int64(rand.Intn(3)+1)) * time.Second)
			newHabitWorker(hw.d, hw.h)
		}
	}()
	for {
		select {
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
		a: Attributes{LastDone: time.Now()},
		h: h,

		cr:  cron.New(),
		cmd: make(chan workerCommand, 10),
		sub: d.Sub(
			daryl.ADD_HABIT_TOPIC,
			fmt.Sprintf("%s.%s", daryl.USER_MESSAGE_TOPIC, h.Id),
		),
	}
	//hw.cr.AddFunc(h.Cron, func() { hw.cmd <- &workerCommandOnHabitTrigger{} })
	hw.cr.AddFunc("0 */10 * * * *", func() { hw.cmd <- &workerCommandTick{} })
	go habitWorkerProcess(hw)
	hw.cr.Start()
	return hw
}

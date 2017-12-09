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

type habitWorker struct {
	d *daryl.Daryl
	a Attributes
	h model.Habit

	cr  *cron.Cron
	cmd chan workerCommand
	sub chan interface{}
}

func (hw *habitWorker) GetHabit() model.Habit {
	r := make(chan model.Habit)
	hw.cmd <- &workerCommandGetHabit{r}
	h := <-r
	close(r)
	return h
}

func (hw *habitWorker) Trigger(t daryl.Trigger) {
}

func (hw *habitWorker) GetWeight() int {
	a := hw.getAttributes()
	return a.Urgent
}

func (hw *habitWorker) getAttributes() Attributes {
	r := make(chan Attributes)
	hw.cmd <- &workerCommandGetAttribute{r}
	a := <-r
	close(r)
	return a
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

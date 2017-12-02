package session

import (
	"errors"
	"time"

	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

type sessionWorkerCommand interface {
	execute(*sessionWorker)
}

type sessionWorkerCommandStop struct {
}

func (sws *sessionWorkerCommandStop) execute(sw *sessionWorker) {
	close(sw.cmd)
}

type sessionWorker struct {
	d   *daryl.Daryl
	r   *protodef.StartWorkSessionRequest
	cmd chan sessionWorkerCommand

	s   model.Session
	due []model.Habit
}

func (sw *sessionWorker) stop() {
}

func sessionWorkerProcess(sw *sessionWorker) {
	for range sw.cmd {
	}
}

func newSessionWorker(d *daryl.Daryl, r *protodef.StartWorkSessionRequest) (*sessionWorker, model.Session, error) {
	_, err := time.ParseDuration(r.Config.Duration)
	if err != nil {
		return nil, model.Session{}, err
	}

	due := d.HabitProcessor.GetDueHabits()

	if len(due) == 0 {
		return nil, model.Session{}, errors.New("All good ! You're free !")
	}

	pss := []model.SessionSlice{}
	for _, d := range due {
		pss = append(pss, model.SessionSlice{
			Start: time.Now(),
			End:   time.Now(),
			Habit: d,
		})
	}
	s := model.Session{
		Start:  time.Now(),
		End:    time.Now(),
		Slices: pss,
	}

	sw := &sessionWorker{d, r, make(chan sessionWorkerCommand), s, due}
	go sessionWorkerProcess(sw)
	return sw, s, nil
}

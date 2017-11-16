package session

import (
	"errors"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type sessionWorkerCommand interface {
	execute(*sessionWorker)
}

type sessionWorker struct {
	d   *daryl.Daryl
	r   *protodef.StartWorkSessionRequest
	cmd chan sessionWorkerCommand

	s   *protodef.Session
	due []*protodef.Habit
}

func (sw *sessionWorker) stop() {
	close(sw.cmd)
}

func sessionWorkerProcess(sw *sessionWorker) {
	for c := range sw.cmd {
		log.Info(c)
	}
}

func newSessionWorker(d *daryl.Daryl, r *protodef.StartWorkSessionRequest) (*sessionWorker, *protodef.Session, error) {
	due := d.HabitProcessor.GetDueHabits()
	log.Info(due)

	if len(due) == 0 {
		return nil, nil, errors.New("All good ! You're free !")
	}

	ss := make([]*protodef.SessionSlice, 0)
	for _, d := range due {
		ss = append(ss, &protodef.SessionSlice{
			Start: ptypes.TimestampNow(),
			End:   ptypes.TimestampNow(),
			Habit: d,
		})
	}
	s := &protodef.Session{
		Start:  ptypes.TimestampNow(),
		End:    ptypes.TimestampNow(),
		Slices: ss,
	}

	sw := &sessionWorker{d, r, make(chan sessionWorkerCommand), s, due}
	go sessionWorkerProcess(sw)
	return sw, s, nil
}

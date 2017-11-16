package session

import (
	"errors"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type session struct {
	s protodef.Session

	slices []daryl.SessionSlice
}

func newSession(ps *protodef.Session) *session {
	s := &session{
		s:      *ps,
		slices: make([]daryl.SessionSlice, 0),
	}
	for _, ss := range ps.Slices {
		s.slices = append(s.slices, newSessionSlice(ss))
	}
	return s
}

func (s *session) GetSession() protodef.Session {
	return s.s
}

func (s *session) GetSessionSlices() []daryl.SessionSlice {
	return s.slices
}

type sessionSlice struct {
	ss protodef.SessionSlice
}

func newSessionSlice(ss *protodef.SessionSlice) *sessionSlice {
	return &sessionSlice{*ss}
}

func (ss *sessionSlice) GetSessionSlice() protodef.SessionSlice {
	return ss.ss
}

type sessionWorkerCommand interface {
	execute(*sessionWorker)
}

type sessionWorker struct {
	d   *daryl.Daryl
	r   *protodef.StartWorkSessionRequest
	cmd chan sessionWorkerCommand

	s   *session
	due []daryl.Habit
}

func (sw *sessionWorker) stop() {
	close(sw.cmd)
}

func sessionWorkerProcess(sw *sessionWorker) {
	for c := range sw.cmd {
		log.Info(c)
	}
}

func newSessionWorker(d *daryl.Daryl, r *protodef.StartWorkSessionRequest) (*sessionWorker, daryl.Session, error) {
	due := d.HabitProcessor.GetDueHabits()

	if len(due) == 0 {
		return nil, nil, errors.New("All good ! You're free !")
	}

	pss := make([]*protodef.SessionSlice, 0)
	for _, d := range due {
		h := d.GetHabit()
		pss = append(pss, &protodef.SessionSlice{
			Start: ptypes.TimestampNow(),
			End:   ptypes.TimestampNow(),
			Habit: &h,
		})
	}
	ps := &protodef.Session{
		Start:  ptypes.TimestampNow(),
		End:    ptypes.TimestampNow(),
		Slices: pss,
	}

	s := newSession(ps)

	sw := &sessionWorker{d, r, make(chan sessionWorkerCommand), s, due}
	go sessionWorkerProcess(sw)
	return sw, s, nil
}

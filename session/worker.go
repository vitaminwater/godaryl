package session

import (
	"errors"
	"sort"

	"github.com/golang/protobuf/ptypes"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type session struct {
	protodef.Session

	slices []daryl.SessionSlice
}

func newSession(ps *protodef.Session) *session {
	s := &session{
		Session: *ps,
		slices:  make([]daryl.SessionSlice, 0),
	}
	for _, ss := range ps.Slices {
		s.slices = append(s.slices, newSessionSlice(ss))
	}
	return s
}

func (s *session) GetSession() protodef.Session {
	return s.Session
}

func (s *session) GetSessionSlices() []daryl.SessionSlice {
	return s.slices
}

type sessionSlice struct {
	protodef.SessionSlice
}

func newSessionSlice(ss *protodef.SessionSlice) *sessionSlice {
	return &sessionSlice{*ss}
}

func (ss *sessionSlice) GetSessionSlice() protodef.SessionSlice {
	return ss.SessionSlice
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
	for _ = range sw.cmd {
	}
}

type sortedHabits []daryl.Habit

func (sh sortedHabits) Len() int {
	return len(sh)
}

func (sh sortedHabits) Less(i, j int) bool {
	return sh[i].GetWeight() < sh[j].GetWeight()
}

func (sh sortedHabits) Swap(i, j int) {
	tmp := sh[j]
	sh[j] = sh[i]
	sh[i] = tmp
}

func newSessionWorker(d *daryl.Daryl, r *protodef.StartWorkSessionRequest) (*sessionWorker, daryl.Session, error) {
	due := sortedHabits(d.HabitProcessor.GetDueHabits())

	if len(due) == 0 {
		return nil, nil, errors.New("All good ! You're free !")
	}

	sort.Sort(due)

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

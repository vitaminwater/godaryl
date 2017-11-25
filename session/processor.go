package session

import (
	"errors"

	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type sessionProcessor struct {
	d  *daryl.Daryl
	sw *sessionWorker
}

func (mp *sessionProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
}

func (sp *sessionProcessor) StartWorkSession(r *protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error) {
	if sp.sw != nil {
		return nil, errors.New("Working session already started, stop current one before starting another one")
	}
	sw, s, err := newSessionWorker(sp.d, r)
	if err != nil {
		return nil, err
	}
	sp.d.Pub(r, daryl.START_WORK_SESSION_TOPIC)
	sp.sw = sw
	sp.d.Pub(s, PROPOSE_WORK_SESSION_TOPIC)
	se := s.GetSession()
	return &protodef.StartWorkSessionResponse{Session: &se}, nil
}

func (sp *sessionProcessor) CancelWorkSession(*protodef.CancelWorkSessionRequest) (*protodef.CancelWorkSessionResponse, error) {
	sp.sw.stop()
	sp.sw = nil
	return &protodef.CancelWorkSessionResponse{}, nil
}

func (sp *sessionProcessor) RefuseSessionSlice(*protodef.RefuseSessionSliceRequest) (*protodef.RefuseSessionSliceResponse, error) {
	if sp.sw == nil {
		return nil, errors.New("No work session yet, create it first")
	}
	return &protodef.RefuseSessionSliceResponse{}, nil
}

func NewSessionProcessor() *sessionProcessor {
	sp := &sessionProcessor{}
	return sp
}

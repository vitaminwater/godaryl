package session

import (
	"github.com/vitaminwater/daryl/daryl"
)

type sessionProcessor struct {
	d *daryl.Daryl
}

func (mp *sessionProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
}

func (sp *sessionProcessor) StartWorkSession(r *daryl.StartWorkSessionRequest) (*daryl.StartWorkSessionResponse, error) {
	sp.d.Pub(r, daryl.START_WORK_SESSION_TOPIC)
	return &daryl.StartWorkSessionResponse{}, nil
}

func NewSessionProcessor() *sessionProcessor {
	sp := &sessionProcessor{}
	return sp
}

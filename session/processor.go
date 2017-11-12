package session

import (
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type sessionProcessor struct {
	d *daryl.Daryl
}

func (mp *sessionProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
}

func (sp *sessionProcessor) StartWorkSession(r *protodef.StartWorkSessionRequest) (*protodef.StartWorkSessionResponse, error) {
	sp.d.Pub(r, daryl.START_WORK_SESSION_TOPIC)
	return &protodef.StartWorkSessionResponse{}, nil
}

func NewSessionProcessor() *sessionProcessor {
	sp := &sessionProcessor{}
	return sp
}

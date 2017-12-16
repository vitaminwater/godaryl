package message

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

type messageTypeProcessor interface {
	process(model.Message)
}

type messageProcessor struct {
	d          *daryl.Daryl
	processors []messageTypeProcessor
	threads    sync.Map
}

func (mp *messageProcessor) SetDaryl(d *daryl.Daryl) {
	mp.d = d
	hs, err := d.HabitProcessor.GetHabits()
	if err != nil {
		log.Warning(err)
	}
	for _, h := range hs {
		id := fmt.Sprintf("habit_%s", h.GetHabit().Id)
		t, err := newThread(id, mp.d, []conversation{habitConversation{h: h}})
		if err != nil {
			log.Warning(err)
		}
		mp.threads.Store(id, t)
	}

	id := fmt.Sprintf("daryl_%s", d.D.Id)
	t, err := newThread(id, mp.d, []conversation{darylConversation{d: d}})
	if err != nil {
		log.Warning(err)
	}
	mp.threads.Store(id, t)
}

func (mp *messageProcessor) UserMessage(r *protodef.UserMessageRequest) (*protodef.UserMessageResponse, error) {
	r.Message.At = ptypes.TimestampNow()
	m, err := model.NewMessageFromProtodef(mp.d.D, r.Message)
	if err != nil {
		return nil, err
	}

	err = m.Insert()
	if err != nil {
		return nil, err
	}

	for _, processor := range mp.processors {
		processor.process(m)
	}

	err = m.Update()
	if err != nil {
		return nil, err
	}

	mp.d.Pub(m, daryl.USER_MESSAGE_TOPIC, fmt.Sprintf("%s.%s", daryl.USER_MESSAGE_TOPIC, m.HabitId.String))

	id := fmt.Sprintf("daryl_%s", mp.d.D.Id)
	if r.Message.HabitIdentifier != "" {
		id = fmt.Sprintf("habit_%s", r.Message.HabitIdentifier)
	}

	t, ok := mp.threads.Load(id)
	if ok == true {
		t.(thread).pushUserMessage(m)
	}

	mm, err := m.ToProtodef()
	if err != nil {
		return nil, err
	}
	return &protodef.UserMessageResponse{Message: mm}, nil
}

func NewMessageProcessor() *messageProcessor {
	mp := &messageProcessor{}
	mp.processors = []messageTypeProcessor{
		newLinkMessageProcessor(mp),
		newTodoMessageProcessor(mp),
	}
	return mp
}

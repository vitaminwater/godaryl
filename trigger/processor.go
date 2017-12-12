package trigger

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

var triggers sync.Map = sync.Map{}

type triggerProcessor struct {
	d *daryl.Daryl
}

func (tp *triggerProcessor) SetDaryl(d *daryl.Daryl) {
	tp.d = d

	ts, err := model.TriggersForDaryl(d.D)
	if err != nil {
		log.Warning(err)
	}
	for _, t := range ts {
		err := tp.addTrigger(t)
		if err != nil {
			log.Warning(err)
		}
	}
}

func (tp *triggerProcessor) AddTrigger(r *protodef.AddTriggerRequest) (*protodef.AddTriggerResponse, error) {
	h, err := tp.d.HabitProcessor.GetHabit(r.Trigger.HabitIdentifier)
	if err != nil {
		return nil, err
	}

	t, err := model.NewTriggerFromProtodef(h.GetHabit(), r.Trigger)
	if err != nil {
		return nil, err
	}

	if t.Id == "" {
		err = t.Insert()
		if err != nil {
			return nil, err
		}
	}

	err = tp.addTrigger(t)
	if err != nil {
		return nil, err
	}

	tpr, err := t.ToProtodef()
	if err != nil {
		return nil, err
	}

	return &protodef.AddTriggerResponse{Trigger: tpr}, nil
}

func (tp *triggerProcessor) IncomingMessage(m *protodef.IncomingTriggerMessageRequest) (*protodef.IncomingTriggerMessageResponse, error) {
	return nil, nil
}

func (tp *triggerProcessor) addTrigger(t model.Trigger) error {
	switch t.Engine {
	case "cron":
		ct, err := newCronTrigger(tp.d, t)
		if err != nil {
			return err
		}
		triggers.Store(t.Id, ct)
	default:
		log.Warningf("Trigger engine %s not found", t.Engine)
	}
	return nil
}

func NewTriggerProcessor() daryl.TriggerProcessor {
	return &triggerProcessor{}
}

package trigger

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

var triggers sync.Map = sync.Map{}

type triggerProcessor struct {
	d *daryl.Daryl
}

func (tp *triggerProcessor) SetDaryl(d *daryl.Daryl) {
	tp.d = d
}

func (hp *triggerProcessor) AddTrigger(r *protodef.AddTriggerRequest) (*protodef.AddTriggerResponse, error) {
	h, err := hp.d.HabitProcessor.GetHabit(r.Trigger.HabitIdentifier)
	if err != nil {
		return nil, err
	}

	t, err := model.NewTriggerFromProtodef(h.GetHabit(), r.Trigger)
	if err != nil {
		return nil, err
	}

	err = daryl_db.Insert("habit_trigger", &t)
	if err != nil {
		return nil, err
	}

	switch r.Trigger.Engine {
	case "cron":
		ct, err := newCronTrigger(hp.d, t)
		if err != nil {
			return nil, err
		}
		triggers.Store(t.Id, ct)
	default:
		log.Warningf("Trigger %s not found", t.Name)
	}

	tp, err := t.ToProtodef()
	if err != nil {
		return nil, err
	}

	return &protodef.AddTriggerResponse{Trigger: tp}, nil
}

func NewTriggerProcessor() daryl.TriggerProcessor {
	return &triggerProcessor{}
}

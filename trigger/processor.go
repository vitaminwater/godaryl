package trigger

import (
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
	"github.com/vitaminwater/daryl/protodef"
)

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

	tp, err := t.ToProtodef()
	if err != nil {
		return nil, err
	}

	return &protodef.AddTriggerResponse{Trigger: tp}, nil
}

func NewTriggerProcessor() daryl.TriggerProcessor {
	return &triggerProcessor{}
}

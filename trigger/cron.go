package trigger

import (
	"github.com/mitchellh/mapstructure"
	"github.com/robfig/cron"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type cronTriggerParams struct {
	Cron string `json:"cron" mapstructure:"cron"`
}

type cronTrigger struct {
	habit  daryl.Habit
	params cronTriggerParams

	t    model.Trigger
	cron *cron.Cron
}

func newCronTrigger(d *daryl.Daryl, t model.Trigger) (cronTrigger, error) {
	h, err := d.HabitProcessor.GetHabit(t.HabitId)
	if err != nil {
		return cronTrigger{}, err
	}

	ct := cronTrigger{habit: h, params: cronTriggerParams{}, t: t}
	mapstructure.Decode(t.Params, &ct.params)

	ct.cron = cron.New()
	ct.cron.AddFunc(ct.params.Cron, func() {
		d.Pub(t, CRON_TRIGGERED_TOPIC)
		h.Trigger(ct)
	})
	ct.cron.Start()
	return ct, err
}

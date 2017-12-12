package trigger

import (
	"github.com/mitchellh/mapstructure"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
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
	log.Info(t.Params)
	mapstructure.Decode(t.Params, &ct.params)
	log.Info(ct.params)

	log.Info(ct.params)
	ct.cron = cron.New()
	ct.cron.AddFunc(ct.params.Cron, func() { h.Trigger(ct) })
	ct.cron.Start()
	return ct, err
}

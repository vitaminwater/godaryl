package habit

import (
	//log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
)

type habitStore struct {
	d      *daryl.Daryl
	c      chan interface{}
	habits []*habitWorker
}

func habitStoreProcess(hs *habitStore) {
	for msg := range hs.c {
		tm := msg.(daryl.TopicMessage)
		switch t := tm.Topic; t {
		case daryl.ADD_HABIT_TOPIC:
			r := tm.Msg.(*protodef.AddHabitRequest)
			h := newHabitWorker(hs.d, r.Habit)
			hs.habits = append(hs.habits, h)
		}
	}
}

func newHabitStore(d *daryl.Daryl) *habitStore {
	hs := &habitStore{d: d, habits: make([]*habitWorker, 0, 10)}
	hs.c = d.Sub(
		daryl.ADD_HABIT_TOPIC,
	)
	go habitStoreProcess(hs)
	return hs
}

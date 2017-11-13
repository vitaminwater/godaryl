package habit

import (
	//log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
)

type command interface {
	execute(hs *habitStore)
}

type commandAddHabit struct {
	h *habit
}

func (c *commandAddHabit) execute(hs *habitStore) {
	hw := newHabitWorker(hs.d, c.h)
	hs.habitWorkers = append(hs.habitWorkers, hw)
}

type habitStore struct {
	d            *daryl.Daryl
	c            chan command
	habitWorkers []*habitWorker
}

func (hs *habitStore) addHabit(h *habit) {
	hs.c <- &commandAddHabit{h}
}

func habitStoreProcess(hs *habitStore) {
	for cmd := range hs.c {
		cmd.execute(hs)
	}
}

func newHabitStore(d *daryl.Daryl) *habitStore {
	hs := &habitStore{d: d, habitWorkers: make([]*habitWorker, 0, 10)}
	go habitStoreProcess(hs)
	return hs
}

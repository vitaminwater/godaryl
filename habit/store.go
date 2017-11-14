package habit

import (
	//log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
)

type storeCommand interface {
	execute(hs *habitStore)
}

type storeCommandAddHabit struct {
	h *habit
}

func (c *storeCommandAddHabit) execute(hs *habitStore) {
	hw := newHabitWorker(hs.d, c.h)
	hs.habitWorkers = append(hs.habitWorkers, hw)
}

type habitStore struct {
	d            *daryl.Daryl
	c            chan storeCommand
	habitWorkers []*habitWorker
}

func (hs *habitStore) addHabit(h *habit) {
	hs.c <- &storeCommandAddHabit{h}
}

func habitStoreProcess(hs *habitStore) {
	for cmd := range hs.c {
		cmd.execute(hs)
	}
}

func newHabitStore(d *daryl.Daryl) *habitStore {
	hs := &habitStore{
		d,
		make(chan storeCommand, 10),
		make([]*habitWorker, 0, 10),
	}
	go habitStoreProcess(hs)
	return hs
}

package habit

import (
	//log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/protodef"
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

type storeCommandGetDueHabit struct {
	r chan []*protodef.Habit
}

func (d *storeCommandGetDueHabit) execute(hs *habitStore) {
	habits := make([]*protodef.Habit, 0, 10)
	for _, w := range hs.habitWorkers {
		r := make(chan protodef.Habit)
		w.cmd <- &workerCommandGetHabit{r}
		h := <-r
		close(r)
		if h.Stats.NMissed > 0 {
			habits = append(habits, &h)
		}
	}
	d.r <- habits
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

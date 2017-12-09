package habit

import (
	//log "github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type storeCommand interface {
	execute(hs *habitStore)
}

type storeCommandGetAttrs struct {
	h model.Habit
	r chan Attributes
}

func (c *storeCommandGetAttrs) execute(hs *habitStore) {
	for _, w := range hs.habitWorkers {
		h := w.GetHabit()
		if h.h.Id == c.h.Id {
			c.r <- h.a
			return
		}
	}
}

type storeCommandAddHabit struct {
	h model.Habit
	r chan model.Habit
}

func (c *storeCommandAddHabit) execute(hs *habitStore) {
	if c.h.Id == "" {
		err := c.h.Insert()
		if err != nil {
			log.Info(err)
			return
		}
	}
	hw := newHabitWorker(hs.d, c.h)
	hs.habitWorkers = append(hs.habitWorkers, hw)
	c.r <- c.h
}

type storeCommandGetDueHabits struct {
	r chan []model.Habit
}

func (d *storeCommandGetDueHabits) execute(hs *habitStore) {
	habits := make([]model.Habit, 0, 10)
	for _, w := range hs.habitWorkers {
		h := w.GetHabit()
		if h.a.NMissed > 0 {
			habits = append(habits, h.h)
		}
	}
	d.r <- habits
}

type habitStore struct {
	d            *daryl.Daryl
	c            chan storeCommand
	habitWorkers []*habitWorker
}

func (s *habitStore) getDueHabits() []model.Habit {
	r := make(chan []model.Habit)
	s.c <- &storeCommandGetDueHabits{r}
	hs := <-r
	close(r)
	return hs
}

func (s *habitStore) getAttributes(h model.Habit) Attributes {
	r := make(chan Attributes)
	s.c <- &storeCommandGetAttrs{h, r}
	a := <-r
	close(r)
	return a
}

func (hs *habitStore) addHabit(h model.Habit) model.Habit {
	r := make(chan model.Habit)
	hs.c <- &storeCommandAddHabit{h: h, r: r}
	h = <-r
	close(r)
	return h
}

func habitStoreProcess(hs *habitStore) {
	for cmd := range hs.c {
		cmd.execute(hs)
	}
}

func (hs *habitStore) loadHabits() error {
	habits, err := model.HabitsForDaryl(hs.d.D)
	if err != nil {
		return err
	}
	for _, h := range habits {
		hs.addHabit(h)
	}
	return nil
}

func newHabitStore(d *daryl.Daryl) *habitStore {
	hs := &habitStore{
		d:            d,
		c:            make(chan storeCommand, 10),
		habitWorkers: make([]*habitWorker, 0, 10),
	}
	go habitStoreProcess(hs)
	if err := hs.loadHabits(); err != nil {
		log.Warning(err)
	}
	return hs
}

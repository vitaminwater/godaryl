package habit

import (
	//log "github.com/sirupsen/logrus"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl/daryl"
	"github.com/vitaminwater/daryl/model"
)

type habitStore struct {
	d            *daryl.Daryl
	c            chan storeCommand
	habitWorkers map[string]*habitWorker
}

func (s *habitStore) getDueHabits() []daryl.Habit {
	r := make(chan []daryl.Habit)
	s.c <- &storeCommandGetDueHabits{r}
	hs := <-r
	close(r)
	return hs
}

func (s *habitStore) getHabit(id string) (daryl.Habit, error) {
	r := make(chan storeCommandGetHabitWorkerResponse)
	s.c <- &storeCommandGetHabitWorker{id: id, r: r}
	hs := <-r
	close(r)
	return hs.hw, hs.err
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
	defer func() {
		if err := recover(); err != nil {
			log.Warn(err)
			time.Sleep(time.Duration(int64(rand.Intn(3)+1)) * time.Second)
			go habitStoreProcess(hs)
		}
	}()

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
		habitWorkers: make(map[string]*habitWorker),
	}
	go habitStoreProcess(hs)
	if err := hs.loadHabits(); err != nil {
		log.Warning(err)
	}
	return hs
}

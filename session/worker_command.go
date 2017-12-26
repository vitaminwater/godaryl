package session

import "github.com/vitaminwater/daryl/model"

type sessionWorkerCommand interface {
	execute(*sessionWorker)
}

type sessionWorkerCommandStop struct {
}

func (sws *sessionWorkerCommandStop) execute(sw *sessionWorker) {
	close(sw.cmd)
}

type getSessionCommand struct {
	r chan model.Session
}

func (sws *getSessionCommand) execute(sw *sessionWorker) {
	sws.r <- sw.s
}

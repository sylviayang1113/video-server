package taskrunner

import (
	"time"
	"log"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTickers(interval *time.Second)
		runner: r,
	}
}

func (w *worker) startWorker()  {
	for {
		select {
		case  <- w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start()  {
	// start videw file clean
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}


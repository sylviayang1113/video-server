package taskrunner

import (

)

type Runner struct {
	Controller controlChan
	Error controlChan
	Data dataChan
	dataSize int
	longLived bool
	Dispatcher fn
	Executor fn
}

func NewRunner(size int, longLived bool, d fn, e fn) *Runner {
	return &Runner {
		Controller: make(chan string, 1),
		Error: make(chan string, 1),
		Data: make(chan interface{}, size),
		longLived: longLived,
		dataSize: size,
		Dispatcher: d,
		Executor: e,
	}
}
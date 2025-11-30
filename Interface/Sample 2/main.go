package main

import "fmt"

type Worker interface {
	Work()
}

type WorkerImpl struct {
	WorkDone bool
}

func (w *WorkerImpl) Work() {
	fmt.Println("I'm working")
	w.WorkDone = true
}

func NewWorker() Worker {
	var w *WorkerImpl = nil
	w = &WorkerImpl{}
	return w
}

func main() {
	var w Worker
	w = NewWorker()
	//_ = w

	if w != nil {
		w.Work()
	}
}

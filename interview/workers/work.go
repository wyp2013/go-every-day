package workers

import "fmt"

type Worker struct {
	workPool chan *Worker
	rev      chan interface{}
	close    chan struct{}
}

func NewWorker(workPool chan *Worker) *Worker {
	return &Worker{
		workPool: workPool,
		close: make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go w.do()
}

func (w *Worker) Stop() {
	close(w.close)
}

func (w *Worker) do() {
	for {
		w.workPool <- w

		select {
		case work, ok := <-w.rev:
			if !ok {
				fmt.Println("rev worker is error")
				return
			}

			if value, ok := work.(int); ok {
				fmt.Println("receive work: ", value)
			}

		case <-w.close:
			fmt.Println("worker receiver close signal, return ")
			return
		}
	}
}

func (w *Worker) Put(work interface{}) {
	w.rev <- work
}


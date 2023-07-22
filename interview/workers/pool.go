package workers

import (
	"context"
	"errors"
	"time"
)

const work_pool_size = 100
type PoolWorker struct {
	WorkPool chan *Worker
	RevChan  chan interface{}
}

func (pw *PoolWorker) Start() {
	pw.WorkPool = make(chan *Worker, work_pool_size)
	pw.RevChan  = make(chan interface{}, 10)

	for i:=0; i < work_pool_size; i++ {
		worker := NewWorker(pw.WorkPool)
		worker.Start()
	}
}

func (pw *PoolWorker) Run() {
	for {
		select {
		case work := <-pw.RevChan:
			go func(work interface{}) {
				worker := <-pw.WorkPool
				worker.Put(work)
			}(work)
		}
	}

}

func (pw *PoolWorker) PutWork(work interface{}) error {
	ctx, _:= context.WithTimeout(context.Background(), time.Duration(5) * time.Second)

	select {
	case pw.RevChan <- work:
		return nil
	case <-ctx.Done():
		return errors.New("time out")
	}
}

package workerpool

import (
	"fmt"
	"sync"
)

const maxJobQueue = 100
const maxWorker   = 10

type JobQueue   chan Job
type WorkerPool chan chan Job
type Dispatcher struct {
	workerPool  WorkerPool
	workers     []*Worker
	quit        chan struct{}
}

var jobQueue JobQueue
func init() {
	jobQueue = make(JobQueue, maxJobQueue)
}

func NewDispatcher() *Dispatcher {
	pool := make(WorkerPool, maxWorker)
	return &Dispatcher{
		workerPool: pool,
		workers:    make([]*Worker, maxWorker),
		quit:       make(chan struct{}),
	}
}

func (d *Dispatcher) Run(jobFunc ProcessJobFunc) {
	for i := 0; i < len(d.workers); i++ {
		d.workers[i] = NewWorker(d.workerPool)
		d.workers[i].Start(jobFunc)
	}

	go d.dispatch()
}

func (d *Dispatcher) Stop() {
	fmt.Println("close dispatcher")
	close(d.quit)

	for i := 0; i < len(d.workers); i++ {
		d.workers[i].Stop()
	}
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-jobQueue:
			// 如果取到任务，就分配给这个worker
			go func(job Job) {
				// 获取一个空闲的worker
				jobChannel := <- d.workerPool

				// 向worker里面写任务
				jobChannel <- job
			}(job)

		case <-d.quit:
			return
		}
	}
}

func PutJob(job Job) {
	jobQueue <- job
}


type Worker struct {
	workerPool WorkerPool
	job        chan Job
	quit       chan struct{}
	once       sync.Once
}

func NewWorker(workerPool WorkerPool) *Worker {
	return &Worker{
		workerPool: workerPool,
		job:        make(chan Job),
		quit:       make(chan struct{}),
	}
}

func (w *Worker) Start(process ProcessJobFunc) {
	go func() {
		for {
			// 注册自己到工作池中，表明当前自己空闲
			w.workerPool <- w.job

			select {
			case job := <-w.job:
				//  取到任务处理
				err := process(job)
				if err != nil {
					// todo
				}

			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.once.Do(func() {
		close(w.quit)
	})
}
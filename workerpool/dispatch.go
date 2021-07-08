package workerpool

const maxJobQueue = 1000
const maxWorker   = 200

var jobQueue = make(JobQueue, maxJobQueue)

type ProcessJobFunc func (JobInterface) error

type Dispatcher struct {
	workerPool WorkerPool
	workers    []*Worker
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		workerPool: make(WorkerPool, maxWorker),
		workers:    make([]*Worker, maxWorker),
	}
}

func (d *Dispatcher) Run(jobFunc ProcessJobFunc) {
	for i := 0; i < len(d.workers); i++ {
		d.workers[i] = NewWorker(d.workerPool)
		d.workers[i].Start(jobFunc)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-jobQueue:
			// 如果取到任务，就分配给这个worker
			go func(job JobInterface) {
				// 获取一个空闲的worker
				workerJob := <- d.workerPool

				// 向worker里面写任务
				workerJob <- job
			}(job)
		}
	}
}


func PutJob(job JobInterface) {
	jobQueue <- job
}
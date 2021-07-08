package workerpool


import "sync"

type JobInterface interface {}
type JobQueue   chan JobInterface
type WorkerPool chan chan JobInterface

type Worker struct {
	workerPool WorkerPool
	job        chan JobInterface
	quit       chan struct{}
	once       sync.Once
}

func NewWorker(workerPool WorkerPool) *Worker {
	return &Worker{
		workerPool: workerPool,
		job:        make(chan JobInterface),
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
					println(err.Error())
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


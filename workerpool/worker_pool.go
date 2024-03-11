package workerpool

import (
	"context"
	"fmt"
	"sync"
)

type Task struct {
	Func func(args ...interface{}) *Result
	Args []interface{}
}

type Result struct {
	Value interface{}
	Err   error
}

type WorkerPool interface {
	Start(ctx context.Context)
	Tasks() chan *Task
	Results() chan *Result
}

type workerPool struct {
	numWorkers int
	tasks      chan *Task
	results    chan *Result
	wg         *sync.WaitGroup
}

var _ WorkerPool = (*workerPool)(nil)

func NewWorkerPool(numWorkers int, bufferSize int) *workerPool {
	return &workerPool{
		numWorkers: numWorkers,
		tasks:      make(chan *Task, bufferSize),
		results:    make(chan *Result, bufferSize),
		wg:         &sync.WaitGroup{},
	}
}

func (wp *workerPool) Start(ctx context.Context) {
	// TODO: implementation
	//
	// Starts numWorkers of goroutines, wait until all jobs are done.
	// Remember to closed the result channel before exit.
	ctx2, cancel2 := context.WithCancel(ctx)
	defer close(wp.Results())
	defer cancel2()
	wp.wg.Add(wp.numWorkers)
	for i := 0; i < wp.numWorkers; i++ {
		go wp.run(ctx2)
	}
	// go func() {
	// 	<-ctx.Done()
	// 	cancel2()
	// }()
	wp.wg.Wait()
}

func (wp *workerPool) Tasks() chan *Task {
	return wp.tasks
}

func (wp *workerPool) Results() chan *Result {
	return wp.results
}

func (wp *workerPool) run(ctx context.Context) {
	// TODO: implementation
	//
	// Keeps fetching task from the task channel, do the task,
	// then makes sure to exit if context is done.
	fmt.Println(1)
	defer wp.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, more := <-wp.Tasks():
			if !more {
				return
			}
			// go func() {
			// 	select {
			// 	case <-ctx.Done():
			// 		return
			// 	}
			// }()
			wp.Results() <- v.Func(v.Args...)
		}
	}
}

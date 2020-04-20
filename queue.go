package taskqueue

import (
	"log"
	"sync"
)

// Queue - a queue for enqueueing Tasks to be processed
type Queue struct {

	// Debug Settings
	tLog   *log.Logger // logging
	tDebug bool        // enable debugging

	// Channals
	tQueueChan chan Task      // Task Channal
	tReadyChan chan chan Task // Ready Task Channals

	// Goroutine synchronization
	tDispatcherSync sync.WaitGroup // Work Dispatcher synchronization
	tWorkersSync    sync.WaitGroup // Workers synchronization

	// Queue Workers
	tWorkers []*Worker

	// Quit Queue
	tQuit chan bool
}

// NewQueue - creates a new Task queue
func NewQueue(nW int) *Queue {

	// workers
	w := make([]*Worker, nW, nW)

	// workers synchronization
	ws := sync.WaitGroup{}

	// Ready Task Channals
	rc := make(chan chan Task, nW)

	// create n Workers
	for i := 0; i < nW; i++ {
		w[i] = NewWorker(rc, ws)
	}

	// return Queue
	return &Queue{
		// Channals
		tQueueChan: make(chan Task),
		tReadyChan: rc,

		// Queue Workers
		tWorkers: w,

		// Goroutine synchronization
		tDispatcherSync: sync.WaitGroup{},
		tWorkersSync:    ws,

		// Quit Queue
		tQuit: make(chan bool),
	}
}

// dispatch workers
func (q *Queue) dispatch() {
	q.tDispatcherSync.Add(1)
	for {
		select {
		case Task := <-q.tQueueChan: // We got something in on our queue
			workerChannel := <-q.tReadyChan // Check out an available worker
			workerChannel <- Task           // Send the request to the channel
		case <-q.tQuit:
			for i := 0; i < len(q.tWorkers); i++ {
				q.tWorkers[i].Stop()
			}
			q.tWorkersSync.Wait()
			q.tDispatcherSync.Done()
			return
		}
	}
}

// Start - starts the worker routines and dispatcher routine
func (q *Queue) Start() {
	for i := 0; i < len(q.tWorkers); i++ {
		q.tWorkers[i].Start() // start workers
	}
	go q.dispatch() // queue dispach
}

func (q *Queue) Stop() {
	q.tQuit <- true          // pass quit flag
	q.tDispatcherSync.Wait() // wait
}

// Submit - adds a new Task to be processed
func (q *Queue) Submit(Task Task) {
	q.tQueueChan <- Task
}

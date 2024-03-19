package main

import (
	"fmt"
	"sync"
	"time"
)

// Way to process the tasks
// Functions to execute the worker pool

// Task definition
type Task struct {
	ID int
}

func (t *Task) Process() {
	fmt.Printf("Processing task %d\n", t.ID)
	time.Sleep(time.Second * 2)
}

type WorkerPool struct {
	Tasks       []Task
	Concurrency int
	taskChannel chan Task
	wg          sync.WaitGroup
}

func (wp *WorkerPool) Worker() {
	for task := range wp.taskChannel {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.taskChannel = make(chan Task, len(wp.Tasks))
	for i := 0; i < wp.Concurrency; i++ {
		go wp.Worker()
	}
	// Send tasks to the task channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.taskChannel <- task
	}
	close(wp.taskChannel)
	wp.wg.Wait()
}

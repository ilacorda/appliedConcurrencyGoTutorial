package main

import "fmt"

func main() {
	tasks := make([]Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = Task{ID: i}
	}

	// create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		Concurrency: 5,
	}

	wp.Run()
	fmt.Println("All tasks are done")
}

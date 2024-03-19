package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// worker function processes tasks and increments the task counter atomically.
func worker(workerID int, tasks <-chan Task, wg *sync.WaitGroup, count *int64) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d is processing task %d\n", workerID, task.ID)
		time.Sleep(time.Second) // Simulating work by sleeping for 1 second

		// Atomically increment the count of processed tasks.
		atomic.AddInt64(count, 1)
	}
}

func main() {
	const numWorkers = 5
	const numTasks = 20

	var wg sync.WaitGroup

	// Create a channel to hold tasks.
	tasks := make(chan Task, numTasks)

	// Use an int64 to hold the count of processed tasks, for atomic operations.
	var count int64

	// Start a fixed number of workers.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg, &count)
	}

	// Submit tasks to the tasks channel.
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}

	close(tasks) // Close the channel to signal to the workers that no more tasks will be submitted.
	wg.Wait()    // Wait for all workers to finish processing.

	fmt.Printf("All tasks have been processed. Total processed tasks: %d\n", count)
}

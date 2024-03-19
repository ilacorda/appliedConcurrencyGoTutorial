package main

import (
	"fmt"
	"time"
)

//The for select pattern in Go is a powerful idiom used in concurrent programming to manage multiple channel operations within a loop,
//enabling a goroutine to wait on multiple communication operations. This pattern combines the for loop and the select statement to create a versatile construct that can handle
//various concurrent tasks, such as sending or receiving values on channels, timeout handling, and non-blocking channel operations.

func main() {
	bufferedChan := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	// for select loop pattern with a buffered channel
	// this will not block the main go routine
	// the main go routine will not wait for the channel to be read
	// the main go routine will not wait for the channel to be written
	for _, s := range chars {
		select {
		case bufferedChan <- s:
		}
	}

	close(bufferedChan)

	for result := range bufferedChan {
		fmt.Println(result)
	}

	// infinite loop with select halted by the done channel

	done := make(chan bool)

	go doWork(done)
	time.Sleep(time.Second * 3)
	close(done)
}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		default:
			fmt.Println("default case going in an infiinite loop")
		}
	}
}


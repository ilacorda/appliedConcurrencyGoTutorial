package main

import (
	"fmt"
	"time"
)

func main() {
	// Go routines example

	// Unless time.Sleep is included it will fork off the go routines and just print "hello" - This is called Fork-join model
	// go routines results will not be in a particular order as they are asynchronous
	go numPrint("1")
	go numPrint("2")
	go numPrint("3")

	// this does not sync with the go routines, we need to properly implement it
	time.Sleep(time.Second * 2)

	fmt.Println("hello")

	// Channels
	myChannel := make(chan string)

	// sending data to the channel
	go func() {
		myChannel <- "some beautiful data sent to the channel"
	}()

	// reading data from the channel
	msg := <-myChannel
	fmt.Print(msg)

	// Select example
	// The select statement is used to choose from multiple send/receive channel operations.
	// The select statement blocks until one of its cases can proceed, then it executes that case.
	// It chooses one at random if multiple are ready.
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "here is"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "the data"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Message from ch1: ", msg1)
	case msg2 := <-ch2:
		fmt.Println("Message from ch2: ", msg2)
	}

	// Deadlock scenario
	ch := make(chan string)
	go greet(ch)                // This goroutine is supposed to send a message to the 'ch' channel.
	time.Sleep(5 * time.Second) // Main goroutine sleeps for 5 seconds, during which 'greet' should ideally send a message.

	fmt.Println("Main ready to receive message")
	// If 'greet' hasn't sent a message before this point, the main goroutine will block here waiting indefinitely for a message.
	greeting := <-ch // Potential deadlock if 'greet' goroutine doesn't send a message on 'ch'.

	// The main goroutine will only proceed past this point if it successfully receives a message from 'ch'.
	time.Sleep(2 * time.Second) // Additional sleep after receiving the message.
	fmt.Println("Greeting received!")
	fmt.Println(greeting)

}

func greet(ch chan string) {

}

func numPrint(num string) {
	fmt.Println(num)
}

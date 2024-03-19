package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go greet(ch)
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready to receive message")
	// receive a message from the channel
	greeting := <-ch
	// sleep and print
	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received!")
	fmt.Println(greeting)
}

func greet(ch chan string) {

}

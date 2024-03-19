package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg) // Pass by reference
	wg.Wait()
	goodbye()

	num := 5
	num.length()

}

func hello(group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println("Hello, world!")
}

func goodbye() {
	fmt.Println("Goodbye, world!")
}

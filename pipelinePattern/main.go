package main

import (
	"fmt"
)

func main() {
	// pipeline pattern
	// pipeline pattern is a series of stages connected by channels, where each stage is a group of goroutines running the same function.
	// Each stage has one or more input channels and one or more output channels. Each stage receives values from its input channels, processes those values, and then sends the resulting values on its output channels.
	// In a well-designed pipeline, the stages execute concurrently, and the data flows through the pipeline with a consistent rate, so that the throughput is governed by the slowest stage in the pipeline.

	nums := []int{2, 3, 4, 7, 1}
	// stage 1
	dataChan := sliceToChannel(nums)
	// stage 2
	finalChannel := square(dataChan)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n
		}
		close(out)
	}()

	return out

}
package main

import (
	"fmt"
	"time"
)

// buffered channels (those with explicitly defined capacities) protect programs from leaks
// in case channel values are never read
// time.After sends after awaiting the execution of a goroutine for the defined time
// this enables us to perform timeout operations, useful for external resources or time-bound
// operations
func timeouts() {
	c1, c2 := make(chan string, 1), make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	// one is not completed on time but we can await its execution finishing after the
	// determined time limit of 1 second
	select {
	case job := <-c1:
		fmt.Println("done", job)
	case <-time.After(1 * time.Second):
		fmt.Println("job not complete on time")
		job := <-c1
		fmt.Println("done after time limit", job)
	}

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "two"
	}()

	// the timeout case wont execute since the job is finished within the stipulated time
	// limit
	select {
	case job := <-c2:
		fmt.Println("done", job)
	case <-time.After(time.Second * 2):
		fmt.Println("this wont run")
	}
}

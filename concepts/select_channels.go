package concepts

import (
	"fmt"
	"time"
)

// select asynchronously awaits for a channel operation to process
// so instead of awaiting for the 3 goroutines to finish one by one, which would be
// blocking in that manner, since we are going step by step, we await the processing
// for all 3 goroutines simultaneously!
// the total execution time here is only ~3 seconds- the longest wait duration for our
// functions since this is all simulatenous
// all 3 sleep events are executed concurrently
func SelectChannels() {
	c1, c2 := make(chan string), make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c3 <- "three"

	}()

	// using default would cause it to be the only processed aspect
	// since everything else needs time to be executed and then we receive
	// the values
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
			// default:
			// 	fmt.Println("all done")
		}
	}

}

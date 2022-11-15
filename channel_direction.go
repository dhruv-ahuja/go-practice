package main

import "fmt"

// takes a receive-only channel and sends the msg to that channel
func pings(ping chan<- string, msg string) {
	ping <- msg
	fmt.Println("sent msg")
}

// now we carry the value of the channel and send it to the other channel
// this time ping is a send only chan and pong a receive only chan
func pongs(ping <-chan string, pong chan<- string) {
	// we first need to store the channel's carried value in a variable
	msg := <-ping
	// now we can safely send it over to the other channel
	pong <- msg
}

func pingPong() {
	ping, pong := make(chan string), make(chan string)
	go pings(ping, "secret")
	go pongs(ping, pong)

	// receiving the channel's value and processing it
	msg := <-pong
	fmt.Println("received msg")
	fmt.Println("msg was: ", msg)
}

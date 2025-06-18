package main

import (
	"fmt"
)

//channels are used to communicate between goroutines
// They allow goroutines to synchronize and share data safely without explicit locks.
// Channels can be buffered or unbuffered, and they provide a way to send and receive

func main() {
	// ch:=make(chan int) // Create a channel to communicate between goroutines
	// ch<-1
	// var i=<-ch // Receive a value from the channel
	// fmt.Println(i) // Print the received value

	//this block causes a deadlock because the channel is unbuffered and no goroutine is ready to receive from it

	ch := make(chan int) // Create an unbuffered channel
	go process(ch)
	for i := 0; i < 10; i++ { 
		fmt.Println(<-ch) // Print the received value
	}

}

func process(ch chan int) {
	defer close(ch) // Ensure the channel is closed when done
	for i:= 0; i < 10; i++ {
		ch <- i // Send values to the channel
		fmt.Println("Sent:", i)
	}
}
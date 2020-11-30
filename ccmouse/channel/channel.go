package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int)  {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int)  chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo()  {
	var channel [10]chan<- int
	
	for i := 0; i < 10; i++ {
		channel[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channel[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channel[i] <- 'A' + i
	}

	time.Sleep(time.Microsecond)
}


func bufferedChannel()  {
	c := make(chan int, 1)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Microsecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main()  {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
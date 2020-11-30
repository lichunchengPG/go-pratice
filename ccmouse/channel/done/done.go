package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id int, w worker)  {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup)  worker {
	w := worker{
		in: make(chan int),
		done: func() {
		wg.Done()
	}}
	go doWork(id, w)
	return w
}

func chanDemo()  {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}


	time.Sleep(time.Microsecond)
}

func main()  {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
}
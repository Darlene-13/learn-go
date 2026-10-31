package main

import (
	"fmt"
	"sync"
)

// Imagine the backend receives 100 payment requests
// We do not want to process them sequentially, so we create workers to process them concurrently.

func worker1(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d received job %s\n", id, job)
	}
}

func main() {

	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1) //
		go worker1(i, jobs, &wg)
	}

	for payment := 1; payment <= 100; payment++ {
		jobs <- payment
	}

	close(jobs)
	wg.Wait()

	fmt.Printf("All payment processed!!")

}

/**
Wait Group is a counting semaphore in go
It is typically used when waiting for a group or go routines or tasks to finish.

When a main go routine starts a task each in a new, and it calls WaitGroup.Go and then wait for all the tasks or go routines to finish by calling WaitGroup.wait.

// WAIT GROUP IN TRACKING TASKS
Wait group may also be used in tracking new tasks without using go to start new go routines by using wg.Add(1) or wg.Done()
*/

/**
var wg sync.WaitGroup

wg.Add(1)
go func(){
defer wg.Done()
task2()
}

wg.Wait()
*/

func worker2(id int, jobs2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job1 := range jobs2 {
		fmt.Printf("Worker %d received job %s\n", id, job1)
	}
}

func main2() {
	jobs := make(chan int, 99)

	var wg sync.WaitGroup

	for i := 1; i <= 9; i++ {
		wg.Add(1)
		go worker2(i, jobs, &wg)
	}

	for processOrder := 1; processOrder <= 99; processOrder++ {
		jobs <- processOrder
	}

	close(jobs)
	wg.Wait()
	fmt.Printf("All process order processed!!")
}

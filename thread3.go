package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

type Task struct {
	id, start, stop int
	resultChannel   chan int
}

// Find primes between start and stop. Write results to resultChannel
func worker(id int, jobChannel chan Task, wg *sync.WaitGroup) {

	log.Printf("worker[%d] started, waiting for job..\n", id)
	job := <-jobChannel

	log.Printf("worker[%d] got job[%d]: start=%d, stop=%d\n", id, job.id, job.start, job.stop)

	if job.start < 3 || job.start%2 == 0 {
		job.resultChannel <- 2
		job.start = 3
	}

	for i := job.start; i <= job.stop; i += 2 {
		k := int(math.Sqrt(float64(i)))
		for j := 3; j <= k; j = j + 2 {
			if i%j == 0 {
				goto notPrime
			}
		}
		job.resultChannel <- i
	notPrime:
	}
	log.Printf("worker[%d] done\n", id)
	wg.Done()
}

func reader(resultChannel chan int, doneChannel chan bool) {
	log.Println("reader started, waiting for results...")
	for {
		prime := <-resultChannel
		if prime == -1 {
			break
		}
		fmt.Printf("%d ", prime)
	}
	fmt.Printf("\n")
	doneChannel <- true
}

func main() {
	// Set timer
	startTime := time.Now()
	// Max prime to find
	const maxPrime int = 1000
	// No of workers
	const poolSize int = 5
	// Size of interval for each worker to check
	chunk := int(maxPrime / poolSize)

	log.Printf("main(): maxPrime=%d, poolSize=%d, chunk=%d\n", maxPrime, poolSize, chunk)

	// Create channels for go routines
	jobChannel := make(chan Task, poolSize)
	resultChannel := make(chan int, 4)
	doneChannel := make(chan bool)

	// Start result collector
	go reader(resultChannel, doneChannel)
	//	time.Sleep(2 * time.Second)

	// Create pool of workers
	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go worker(i+1, jobChannel, &wg)
	}
	log.Printf("main(): %d workers created.\n", poolSize)
	//	time.Sleep(2 * time.Second)

	// Send assignments to workers
	jobid := 1
	for i := 1; i < 1000; i += chunk {
		start := i
		stop := i + chunk - 1
		jobChannel <- Task{jobid, start, stop, resultChannel}
		jobid++
	}
	close(jobChannel)

	// Wait for everybody to finish.
	log.Println("main(): waiting...")
	wg.Wait()

	// Send EOF on resultChannel and wait for ack
	resultChannel <- -1
	<-doneChannel
	close(resultChannel)
	close(doneChannel)

	log.Println("main(): ending")
	log.Printf("Elapsed: %.6f seconds\n", time.Now().Sub(startTime).Seconds())
}

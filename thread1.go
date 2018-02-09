package main

import (
	"fmt"
	"time"
)

func hello(p chan int) {
	fmt.Println("-> hello()")
	time.Sleep(1 * time.Second)
	data := <-p
	fmt.Printf("\nhello(): received %d\n", data)
	p <- 1
	fmt.Println("<- hello()")
}

func numbers(p chan int) {
	fmt.Println("-> numbers()")
	data := <-p
	fmt.Printf("\nnumbers(): received %d\n", data)
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	p <- 2
	fmt.Println("<- numbers()")
}
func alphabets(p chan int) {
	fmt.Println("-> alphabets()")
	data := <-p
	fmt.Printf("\nalphabets(): received %d\n", data)
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	p <- 3
	fmt.Println("<- alphabets()")
}

func main() {

	// Create channel
	pipe := make(chan int)

	// Start all threads
	fmt.Println("main(): starting threads..")
	go hello(pipe)
	go numbers(pipe)
	go alphabets(pipe)

	fmt.Println("main(): sleeping for 3 secs")
	time.Sleep(3000 * time.Millisecond)

	// Write to channel
	fmt.Println("Writing to channel x3")
	pipe <- 1
	pipe <- 2
	pipe <- 3

	// Read from channel
	for i := 0; i < 3; i++ {
		data := <-pipe
		fmt.Printf("\nmain(): read %d\n", data)
	}

	// Wait...
	//	time.Sleep(3000 * time.Millisecond)
	//	fmt.Println("main() wake up")

}

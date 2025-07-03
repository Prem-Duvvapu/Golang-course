package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(myCh)
		// ch <- 5
		// ch <- 6
		
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func dataRacer() {
	const raceCars = 100
	counter := 0

	wg.Add(raceCars)

	for i := 0; i < raceCars; i++ {
		go func() {
			raceCar := counter
			runtime.Gosched()

			raceCar++
			counter = raceCar
			fmt.Println("Number of go routines:", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Number of go routines:", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}

func mutexVsAliens() {
	var mutex sync.Mutex

	const raceCars = 100
	counter := 0

	wg.Add(raceCars)

	for i := 0; i < raceCars; i++ {
		go func() {
			mutex.Lock()
			raceCar := counter
			runtime.Gosched()

			raceCar++
			counter = raceCar
			fmt.Println("Number of go routines:", runtime.NumGoroutine())
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Number of go routines:", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}

func atomicPackageUser() {

	const raceCars = 100
	var counter int64

	wg.Add(raceCars)

	for i := 0; i < raceCars; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter is:", atomic.LoadInt64(&counter))
			// runtime.Gosched()
			fmt.Println("Number of go routines:", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Number of go routines:", runtime.NumGoroutine())
	fmt.Println("Counter is", counter)
}

func useChannels() {
	c := make(chan int)

	go send(c)
	receive(c, "")

	fmt.Println("About to exit")
}

func send(c chan<- int) {
	for i := 14; i <= 113; i *= 2 {
		c <- i
	}
	close(c)
}

func receive(c <-chan int, message string) {
	for v := range c {
		fmt.Println(message + fmt.Sprint(v))
	}
	wg.Done()
}

func receiveMultipleEvenOdd(even, odd, quit chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Taken from the even channel:", v)
		case v := <-odd:
			fmt.Println("Taken from the odd channel:", v)
		case v := <-quit:
			fmt.Println("Taken from the quit channel", v)
			return
		}
	}
}

func sendMultiple(channelOne, channelTwo chan int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			channelOne <- i
		} else {
			channelTwo <- i
		}
	}

	close(channelOne)
	close(channelTwo)
}

func selectStatementExercise() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i <= 177; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}

		quit <- 0
	}()

	receiveMultipleEvenOdd(even, odd, quit)

	fmt.Println("About to quit")
}

func fanIn(channelOne, channelTwo chan int) <-chan int {
	returnChannel := make(chan int)

	go func() {
		for {
			returnChannel <- <-channelOne
		}
	}()

	go func() {
		for {
			returnChannel <- <-channelTwo
		}
	}()

	return returnChannel
}

func receiveMultiple(channelOne chan int, messageOne string, channelTwo chan int, messageTwo string) {
	wg.Add(2)
	go receive(channelOne, messageOne)
	go receive(channelTwo, messageTwo)
}

func main() {
	fmt.Println("Channels file is working!")

	channelOne := make(chan int)
	channelTwo := make(chan int)

	quit := make(chan int)
	quit <- 0

	go sendMultiple(channelOne, channelTwo)
	fannedInChannel := fanIn(channelOne, channelTwo)

	receive(fannedInChannel, "Fanned in channel value is: ")

	wg.Wait()
}

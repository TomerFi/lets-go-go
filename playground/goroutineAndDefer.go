package playground

// the following imports are from go's own packages
import (
	"fmt"
	"sync"
	"time"
)

/* ########################### ##
## #### Utility Functions #### ##
## ########################### */

// sleep for x seconds and then print the text
func waitAndPrint(text string, seconds time.Duration) {
	time.Sleep(seconds * time.Second)
	fmt.Println(text)
}

// sleep for x seconds, print the text, and inform the wait group on done
func waitAndPrintWithWaitGroup(text string, seconds time.Duration, waitGroup *sync.WaitGroup) {
	time.Sleep(seconds * time.Second)
	fmt.Println(text)
	waitGroup.Done()
}

// sleep for x seconds and send the text back to the channel
func waitAndReturnTextWithChannel(text string, seconds time.Duration, channel chan<- string) {
	time.Sleep(seconds * time.Second)
	channel <- text
}

/* ############################## ##
## #### Playground Functions #### ##
## ############################## */

func RunDeferAndGoroutine() {
	fmt.Println()
	// defer will run the waitAndPrint function when the runDeferAndGoroutine is done.
	defer waitAndPrint("will be printed when done", 0)
	// this will launch a lightweight thread that will print in one second
	go waitAndPrint("will print in 1 second from a lightweight thread (non-blocking)", 1)
	// this will block the runDeferAndGoroutine for 2 seconds, in that time the goroutine will be done too
	waitAndPrint("will print from this thread in 2 seconds (blocking)", 2)
}

func RunDeferAndGoroutineWithWaitGroup() {
	fmt.Println()
	// create a wait group
	var waitGroup sync.WaitGroup
	// defer will run the waitAndPrint function when the runDeferAndGoroutineWithWaitGroup is done.
	defer waitAndPrint("will be printed when done", 0)
	// this will launch a lightweight thread that will print in two seconds
	go waitAndPrintWithWaitGroup("will print in 2 second from a lightweight thread (non-blocking)", 2, &waitGroup)
	waitGroup.Add(1) // tell the wait group to wait for 1 done notifications
	// this will block the runDeferAndGoroutineWithWaitGroup
	waitAndPrint("will print from this thread now (blocking)", 0)
	// wait for the wait group
	waitGroup.Wait()
}

func RunDeferAndGoroutineWithChannel() {
	fmt.Println()
	// create a string channel
	channel := make(chan string)
	// defer will run the waitAndPrint function when the runDeferAndGoroutineWithChannel is done.
	defer waitAndPrint("will be printed when done", 0)
	// this will launch a lightweight thread that will print in two seconds
	go waitAndReturnTextWithChannel("will send the text back on the channel from a lightweight thread in 2 second (non-blocking)", 2, channel)
	// this will block the runDeferAndGoroutineWithChannel
	waitAndPrint("will print from this thread in 1 second (blocking)", 1)
	// print text from channel now (blocking)
	waitAndPrint(<-channel, 0) // channel is like queue, as long as its open you can feed it and consume from it
}

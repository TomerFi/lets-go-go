package playground

// the following imports are from go's own packages
import (
	"context"
	"fmt"
	"time"
)

/* ########################### ##
## #### Utility Functions #### ##
## ########################### */

// return ping for pong and vice-versa thru the bidirectional channel every 0.5 seconds as long as the context is alive
// the channel is implicitly bidirectional by the lack of the <- operator (see next function)
func pingPongBidirectional(ctx context.Context, bidirectionalChannel chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case gotValue := <-bidirectionalChannel:
			if gotValue == "ping" {
				fmt.Println("ping ->")
				bidirectionalChannel <- "pong"
			} else if gotValue == "pong" {
				fmt.Println("<- pong")
				bidirectionalChannel <- "ping"
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// return ping for pong and vice-versa thru the reading/writing channels every 0.5 seconds as long as the context is alive
// note the <- to the left of the chan type makes it reading only channel, while on the right side makes it a writing only one
func pingPongOneSided(ctx context.Context, readingChannel <-chan string, writingChannel chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		case gotValue := <-readingChannel:
			if gotValue == "ping" {
				fmt.Println("ping ->")
				writingChannel <- "pong"
			} else if gotValue == "pong" {
				fmt.Println("<- pong")
				writingChannel <- "ping"
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}

/* ############################## ##
## #### Playground Functions #### ##
## ############################## */

func GoroutinesPingPongWithBidirectionalChannelManagedByContext() {
	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("Create the parent context at the edge of the application")
	parentContext := context.Background()

	fmt.Println("Create a timeout child context with 6 seconds timeout and gain access for canceling and timing the context")
	timedContext, cancelFunc := context.WithTimeout(parentContext, 6*time.Second)

	fmt.Println("Defer a context cancelation as a best practice")
	defer cancelFunc()

	fmt.Println("Create a string channel")
	bidirectionalChannel := make(chan string)

	fmt.Println("Start two goroutines with the timed context and string channel")
	go pingPongBidirectional(timedContext, bidirectionalChannel)
	go pingPongBidirectional(timedContext, bidirectionalChannel)

	fmt.Println("Send the first ping to get the ball rolling,")
	fmt.Println("and sleep for 7 seconds to the operation will have its 6 seconds to be over")
	bidirectionalChannel <- "ping"
	time.Sleep(7 * time.Second)
}

func GoroutinesPingPongWithOneSidedChannelsManagedByContext() {
	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("Create the parent context at the edge of the application")
	parentContext := context.Background()

	fmt.Println("Create a timeout child context with 6 seconds timeout and gain access for canceling and timing the context")
	timedContext, cancelFunc := context.WithTimeout(parentContext, 6*time.Second)

	fmt.Println("Defer a context cancelation as a best practice")
	defer cancelFunc()

	fmt.Println("Create a string channel for the pings")
	pingsChannel := make(chan string)

	fmt.Println("Create a string channel for the pongs")
	pongsChannel := make(chan string)

	fmt.Println("Start two goroutines with the timed context and string channel reversing the channels between them")
	go pingPongOneSided(timedContext, pingsChannel, pongsChannel)
	go pingPongOneSided(timedContext, pongsChannel, pingsChannel)

	fmt.Println("Send the first ping to get the ball rolling,")
	fmt.Println("and sleep for 7 seconds to the operation will have its 6 seconds to be over")
	pingsChannel <- "ping"
	time.Sleep(7 * time.Second)
}

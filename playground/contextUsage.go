package playground

// the following imports are from go's own packages
import (
	"context"
	"fmt"
	"time"
)

// create a key type for using as a key interface for the context's values
type myKeyType string

/* ########################### ##
## #### Utility Functions #### ##
## ########################### */

// function that takes a context, extract a value by a key and prints it
func functionUsingContext(ctx context.Context, theKey myKeyType) {
	value := ctx.Value(theKey)
	fmt.Printf("Got by key from context: %s", value)
	fmt.Println()
}

// loop until the context is done printing the second every second identifying the invocation with funcName
func loopWhileContextAlive(ctx context.Context, funcName string) {
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err != nil {
				// print the context close reason
				fmt.Printf("Function named %s got reason for ending context: %s", funcName, err)
				fmt.Println()
			}
			return
		default:
			fmt.Printf("Function named %s is inside loop: %d", funcName, time.Now().Second())
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}

// extract the value from the context based on the key an propagate it to the underlying loop function
func loopWhileContextAliveExtractKey(ctx context.Context, theKey myKeyType) {
	loopWhileContextAlive(ctx, ctx.Value(theKey).(string))
}

/* ############################## ##
## #### Playground Functions #### ##
## ############################## */

func SimpleContextExample() {
	fmt.Println()

	fmt.Println("Create the parent context at the edge of the application")
	parentContext := context.Background()

	fmt.Println("Create a with-cancel child context to gain access for canceling all children")
	childCancelContext, cancelFunc := context.WithCancel(parentContext)

	fmt.Println("Defer a context cancelation as a best practice")
	defer cancelFunc()

	fmt.Println("Create a key implementation")
	myKey := myKeyType("myKey")

	fmt.Println("Create a child value context from the with-cancel context")
	childValueCtx := context.WithValue(childCancelContext, myKey, "myValue")

	fmt.Println("Invoke the function using the child value context and expect the value to printed")
	functionUsingContext(childValueCtx, myKey)

	time.Sleep(time.Second)
}

func TimedContextExample() {
	fmt.Println()

	fmt.Println("Create the parent context at the edge of the application")
	parentContext := context.Background()

	fmt.Println("Create a timeout child context with 4 seconds timeout and gain access for canceling and timing the context")
	timedContext, cancelFunc := context.WithTimeout(parentContext, 4*time.Second)

	fmt.Println("Defer a context cancelation as a best practice")
	defer cancelFunc()

	fmt.Println("Invoke the looping function with the child context and expect 4 iterations printed,")
	fmt.Println("also expect the cause for ending the context to be 'context deadline exceeded'")
	loopWhileContextAlive(timedContext, "default-name")

	time.Sleep(time.Second)
}

func ComplexContextExample() {
	fmt.Println()

	fmt.Println("Create the parent context at the edge of the application")
	parentContext := context.Background()

	fmt.Println("Create a timeout child context with 4 seconds timeout and gain access for canceling and timing the context")
	timedContext, cancelFunc := context.WithTimeout(parentContext, 4*time.Second)

	fmt.Println("Defer a context cancelation as a best practice")
	defer cancelFunc()

	fmt.Println("Create a key implementation")
	nameKey := myKeyType("myNameIs")

	fmt.Println("Create two children value-contexts from the timeout context")
	firstChildValueCtx := context.WithValue(timedContext, nameKey, "First child")
	secondChildValueCtx := context.WithValue(timedContext, nameKey, "Second child")

	fmt.Println("Invoke the looping function twice concurrently, once for each context and expect a total of 8 iterations printed,")
	fmt.Println("also expect the cause for ending the both contexts to be 'context deadline exceeded', exemplifying the context hierarchy")
	go loopWhileContextAliveExtractKey(firstChildValueCtx, nameKey)
	go loopWhileContextAliveExtractKey(secondChildValueCtx, nameKey)

	fmt.Println("Sleep for 5 seconds to the operation will have its 4 seconds to be over")
	time.Sleep(5 * time.Second)
}

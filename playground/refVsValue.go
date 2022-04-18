package playground

// and import list instead of multiple imports
import (
	"fmt"               // this is a module provided by go
	"lets-go-go/greets" // note the module name prior to the package name instead of just the relative package name
	"time"              // this is a module provided by go
)

/* ########################### ##
## #### Utility Functions #### ##
## ########################### */

// can take a reference or a value, but will treat it as a value
func functionThatTakesGreeterValue(greeter greets.Greeter) {
	greeter.Greeting = "modifyFromTakesGreeterValue"
}

// the * makes this take a reference only
func functionThatTakesGreeterReference(greeter *greets.Greeter) {
	greeter.Greeting = "modifyFromTakesGreeterReference"
}

/* ############################## ##
## #### Playground Functions #### ##
## ############################## */

func ReferenceVsValues() {
	fmt.Println()
	fmt.Println("Initialize a Greeter with \"originalGreeterValue\" and store the value of it in 'greeterValue'")
	greeterValue := greets.Greeter{
		Greeting: "originalGreeterValue",
	}
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing the stored value as a value to a function that will modify it,")
	functionThatTakesGreeterValue(greeterValue)
	fmt.Println("the modification will, of course, not be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterValue.Greeting)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing a reference of the stored valued to a function that will modify it,")
	functionThatTakesGreeterReference(&greeterValue) // note the & which passes a reference for the greeterValue
	fmt.Println("the modification will, of course, be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterValue.Greeting)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ###########################################################################################################

	fmt.Println()
	fmt.Println("Initialize a Greeter with \"originalGreeterReference\" and store reference of it in 'greeterReference'")
	greeterReference := &greets.Greeter{ // note the & which stores a reference for the initialized Greeter
		Greeting: "originalGreeterReference",
	}
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing the stored reference as a value to a function that will modify it,")
	// note that althugh we are using * to pass the reference as a value,
	// the invoked function treats it as a value
	functionThatTakesGreeterValue(*greeterReference)
	fmt.Println("the modification will, of course, not be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterReference.Greeting)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing the stored reference to a function that will modify it,")
	functionThatTakesGreeterReference(greeterReference)
	fmt.Println("the modification will, of course, be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterReference.Greeting)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}

package playground

// an import list instead of multiple imports
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
	// as this function takes a value, you will never see this modification outside of this function
	greeter.Greeting = "modifiedByFunctionThatTakesGreeterValue"
}

// the * makes this take a reference only
func functionThatTakesGreeterReference(greeter *greets.Greeter) {
	greeter.Greeting = "modifiedByFunctionThatTakesGreeterReference"
}

/* ############################## ##
## #### Playground Functions #### ##
## ############################## */

func ReferenceVsValues() {
	fmt.Println()
	fmt.Println("Initialize a Greeter with \"originalGreeterValue\" and store the value of it in variable named 'greeterValue`")
	greeterValue := greets.Greeter{
		Greeting: "originalGreeterValue",
	}
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing greeterValue as a value to a function that takes a value and modifies it,")
	functionThatTakesGreeterValue(greeterValue)
	fmt.Println("and the modification will not be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterValue.Greeting) // expect originalGreeterValue
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing a reference of the greeterValue variable to a function that takes a reference and modifies it,")
	functionThatTakesGreeterReference(&greeterValue) // note the & which passes a reference for the greeterValue
	fmt.Println("and the modification will be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterValue.Greeting) // expect modifiedByFunctionThatTakesGreeterReference
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ###########################################################################################################

	fmt.Println()
	fmt.Println("Initialize a Greeter with \"originalGreeterReference\" and store a reference of it in a variable named 'greeterReference'")
	greeterReference := &greets.Greeter{ // note the & which stores a reference for the initialized Greeter
		Greeting: "originalGreeterReference",
	}
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing greeterReference as a value to a function that takes a value and modifies it,")
	// note that although we are using * to pass the reference as a value,
	// the invoked function treats it as a value
	functionThatTakesGreeterValue(*greeterReference)
	fmt.Println("and the modification will not be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterReference.Greeting) // expect originalGreeterReference
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Passing greeterReference as a reference to a function that takes a reference and modifies it,")
	functionThatTakesGreeterReference(greeterReference)
	fmt.Println("and the modification will be visible outside that function:")
	fmt.Printf("greeterValue.Greeting = '%s'", greeterReference.Greeting) // expect modifiedByFunctionThatTakesGreeterReference
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}

package playground

// and import list instead of multiple imports
import (
	"fmt"               // this is a module provided by go
	"lets-go-go/greets" // note the module name prior to the package name instead of just the relative package name
	"time"              // this is a module provided by go
)

/* ############################## ##
## #### Playground Functions #### ##
## ############################## */

func FunctionsVsReceiverFunctions() {
	fmt.Println()
	fmt.Println("Initialize a new Greeter struct found in greets/afile with 'Hello' as the Greeting field")
	myGreeter := greets.Greeter{
		Greeting: "Hello",
	}

	fmt.Println()
	fmt.Println("Use the receiver function from the greets/afile to greet the value 'World':")
	fmt.Println(myGreeter.Greet("World"))
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Use the receiver function from the greets/anotherfile,")
	fmt.Println("which uses the (non-receiver) private function from greets/afile to greet again with the value 'World':")
	fmt.Println(myGreeter.GreetAgain("World"))
	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	fmt.Println("Use the (non-receiver) public function from greets/anotherfile to say goodbye to the value 'World':")
	fmt.Println(greets.PublicFunctionSayGoodbye("World"))
	time.Sleep(500 * time.Millisecond)
}

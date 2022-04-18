package greets

import "fmt"

// this is the Greeter struct (a collection of fields)
type Greeter struct {
	// public field of type string (starts with uppercase)
	// this can also be positional fields, ie no names.
	Greeting string
}

// this is a public receiver function (starts with uppercase) referencing the Greeter struct
// takes a string name and returns a string value chained from the Sprintf utility function
func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("%s %s?", g.Greeting, name)
}

// this is a private non-receiver function (starts with lowercase) only accessible within the greets package
// takes a string name and return a string initialized in the function signature
func privateGreetAgainFunction(greeting string, name string) (retValue string) {
	retValue = fmt.Sprintf("I said %s %s??", greeting, name)
	return retValue
}

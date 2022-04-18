package greets

import "fmt"

// this is a public receiver function (starts with uppercase) referencing the Greeter struct from the greets/afile
// takes a string name and returns a string declared and assign using the var keyword
func (g *Greeter) GreetAgain(name string) string {
	// using the private function from greets/afile
	var retValue = privateGreetAgainFunction(g.Greeting, name)
	return retValue
}

// this is a public function (starts with uppercase)
// takes a string name and return a string declared an assigned using the short declaration
func PublicFunctionSayGoodbye(name string) string {
	retValue := fmt.Sprintf("Goodbye %s!", name)
	return retValue
}

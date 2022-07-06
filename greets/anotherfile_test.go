package greets

import (
	"testing"
)

func TestPublicFunctionSuccess(t *testing.T) {
	// invoke the function with a custom value
	retValue := PublicFunctionSayGoodbye("fake name")
	// verify the response and fail accordingly
	if retValue != "Goodbye fake name!" {
		t.Error("this will never happen")
	}
}

func TestPublicFunctionFailure(t *testing.T) {
	// invoke the function with a custom value
	retValue := PublicFunctionSayGoodbye("fake name")
	// THIS WILL FAIL
	// replace this line with ' if retValue != "Goodbye fake name!" ' to make this test pass
	if retValue != "this fails intentionally" {
		t.Error(`the PublicFunctionSayGoodbye function is not working as expected,
		replace "this fails intentionally" with "Goodbye fake name!" to resolve this.`)
	}
}

func TestGreeterAddedReceiverMethod(t *testing.T) {
	// create a value of type Greeter from greets/afile.go
	sut := Greeter{
		Greeting: "yo! wassup",
	}
	// invoke the receiver function Greeter.GreetAgain from greets/anotherfile.go
	gotValue := sut.GreetAgain("dog")
	// verify the response and fail accordingly
	if gotValue != "I said yo! wassup dog??" {
		t.Error("maybe there's something wrong with the private function from afile")
	}
}

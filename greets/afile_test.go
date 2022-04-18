package greets

import "testing"

/*
-   test functions starts with TestXxx (note the camelcase) and takes a reference to the testing.T type
	and reside in files prefixed with _test

-   the test files resides in the same package of the subjects they are testing
	these file's content will not be included in the final binaries

-   the standard convention in regards to testing in Go is that there is no need for assertions,
	a successful test case is a test case that didn't fail/threw an error.
	that being said, there are alternatives: https://github.com/stretchr/testify/

-	to test verbosely and recursively use `go test -v ./...`

-	docs: https://pkg.go.dev/testing
*/
func TestOriginalReceiverFunction(t *testing.T) {
	// create a value of type Greeter from greets/afile.go
	sut := Greeter{
		Greeting: "sup",
	}
	// invoke the receiver function Greeter.Greet from greets/afile.go
	gotValue := sut.Greet("dog")
	// verify the response and fail accordingly
	if gotValue != "sup dog?" {
		t.Error("huston we have a problem")
	}
}

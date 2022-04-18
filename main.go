// main package
package main

/*
go dependency management is both centralized and decentralized.
the packages resides in any git repository,
but will go through the default public proxy https://proxy.golang.org (configurable)
if it's not available in the proxy, it will be dowloaded by it.

for instance, you might see stuff like:
import "github.com/operator-framework/operator-lib"

if this is not already available in your local cache,
go will look for it in default proxy or the one set in the GOPROXY env var.
if not found it will be dowloaded from github.com.

the https://pkg.go.dev/ site recursively scans the public proxy in order to be
the central source of information about go modules.

it's important to note that there are not relative imports in go,
every import require the fully qualified name (more on that in other files).

read: https://go.dev/doc/modules/managing-dependencies
*/

import (
	"fmt"                   // this is a module provided by go
	"lets-go-go/playground" // note the module name prior to the package name instead of just the relative package name
)

/*
the main function is the function being invoked from the main package to start the application
you can invoke it from the command line using `go run .` from the project's root.

you can build for your operating system using `go build .` from the projects' root,
that will create an executable for you to run.
*/
func main() {
	// basic
	playground.FunctionsVsReceiverFunctions()
	playground.ReferenceVsValues()
	// concurency
	playground.RunDeferAndGoroutine()
	playground.RunDeferAndGoroutineWithWaitGroup()
	playground.RunDeferAndGoroutineWithChannel()
	// context
	playground.SimpleContextExample()
	playground.TimedContextExample()
	playground.ComplexContextExample()
	// mix
	playground.GoroutinesPingPongWithBidirectionalChannelManagedByContext()
	playground.GoroutinesPingPongWithOneSidedChannelsManagedByContext()
	// various
	playground.AtomicIntegerPlayground()
	playground.RunTwoGoroutinesIncrementingTheSameAtomicInteger()

	fmt.Println()
}

// helpful go tours can be found here: https://go.dev/tour/list

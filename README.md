# Playground for starting Go development

Testing, functions, receiver functions, and visibility are demonstrated in the [greets package][0].</br>
Go through the code files and follow the comments and log prints.

> Missing playgrounds for the basic stuff like types, arrays, slices, maps.

## Playgrounds

- Basic
  - [Functions VS Receiver Functions][1]
  - [References VS Values][2]
- Concurrency
  - [Simple Defer and Goroutine][3]
  - [Goroutine coordinated with a WaitGroup][4]
  - [Goroutine communicates with a Channel][5]
- Context
  - [Simple Context Usage][6]
  - [Timed Context Usage][7]
  - [Complex Context Usage][8]
- Mix
  - [Two Goroutines playing ping pong with a bidirectional Channel and manged by a Context][9]
  - [Two Goroutines playing ping pong with a two one-sided Channels and manged by a Context][10]
- Various
  - [A partial implantation of Java's AtomicInteger in Go][11]
  - [A demo for using the AtomicInteger][12]
  - [Updating the AtomicInteger concurrently with two Goroutines][13]

<!-- REFS -->
[0]: https://github.com/TomerFi/lets-go-go/tree/main/greets
[1]: https://github.com/TomerFi/lets-go-go/blob/main/playground/funcVsRecFunc.go#L14
[2]: https://github.com/TomerFi/lets-go-go/blob/main/playground/refVsValue.go#L28
[3]: https://github.com/TomerFi/lets-go-go/blob/main/playground/goroutineAndDefer.go#L37
[4]: https://github.com/TomerFi/lets-go-go/blob/main/playground/goroutineAndDefer.go#L47
[5]: https://github.com/TomerFi/lets-go-go/blob/main/playground/goroutineAndDefer.go#L62
[6]: https://github.com/TomerFi/lets-go-go/blob/main/playground/contextUsage.go#L53
[7]: https://github.com/TomerFi/lets-go-go/blob/main/playground/contextUsage.go#L77
[8]: https://github.com/TomerFi/lets-go-go/blob/main/playground/contextUsage.go#L96
[9]: https://github.com/TomerFi/lets-go-go/blob/main/playground/mix.go#L58
[10]: https://github.com/TomerFi/lets-go-go/blob/main/playground/mix.go#L84
[11]: https://github.com/TomerFi/lets-go-go/blob/main/playground/atomicInteger.go#L30
[12]: https://github.com/TomerFi/lets-go-go/blob/main/playground/atomicInteger.go#L116
[13]: https://github.com/TomerFi/lets-go-go/blob/main/playground/atomicInteger.go#L162

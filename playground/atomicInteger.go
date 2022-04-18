package playground

// the following imports are from go's own packages
import (
	"fmt"
	"sync"
	"time"
)

/* ########################### ##
## #### Utility Functions #### ##
## ########################### */

// invoke the GetAndIncrement on the sut for i amount of times waiting to ms milliseconds after each invocation
func incrementBy(sut *AtomicInteger, i int, ms time.Duration, name string) {
	for run := 1; run <= i; run++ {
		fmt.Printf("%s run #%d", name, run)
		fmt.Println()
		current := sut.GetAndIncrement()
		fmt.Printf(" - got %d, incrementing and sleeping for %d ms", current, ms)
		fmt.Println()
		time.Sleep(ms * time.Millisecond)
	}
}

/* ################################################ ##
## #### Implementation of Java's AtomicInteger #### ##
## ################################################ */

type AtomicInteger struct {
	lock  sync.Mutex
	count int
}

func (c *AtomicInteger) Get() int {
	return c.count
}

func (c *AtomicInteger) GetAndSet(newValue int) int {
	defer func() {
		defer c.lock.Unlock()
		c.lock.Lock()
		c.count = newValue
	}()
	return c.count
}

func (c *AtomicInteger) GetAndAdd(delta int) int {
	defer func() {
		defer c.lock.Unlock()
		c.lock.Lock()
		c.count += delta
	}()
	return c.count
}

func (c *AtomicInteger) AddAndGet(delta int) int {
	c.lock.Lock()
	c.count += delta
	c.lock.Unlock()
	return c.count
}

func (c *AtomicInteger) GetAndDecrement() int {
	defer func() {
		defer c.lock.Unlock()
		c.lock.Lock()
		c.count--
	}()
	return c.count
}

func (c *AtomicInteger) DecrementAndGet() int {
	c.lock.Lock()
	c.count--
	c.lock.Unlock()
	return c.count
}

func (c *AtomicInteger) GetAndIncrement() int {
	defer func() {
		defer c.lock.Unlock()
		c.lock.Lock()
		c.count++
	}()
	return c.count
}

func (c *AtomicInteger) IncrementAndGet() int {
	c.lock.Lock()
	c.count++
	c.lock.Unlock()
	return c.count
}

func (c *AtomicInteger) GetAndUpdate(updateFunction func(v int) int) int {
	defer func() {
		defer c.lock.Unlock()
		c.lock.Lock()
		c.count = updateFunction(c.count)
	}()
	return c.count
}

func (c *AtomicInteger) UpdateAndGet(updateFunction func(v int) int) int {
	c.lock.Lock()
	c.count = updateFunction(c.count)
	c.lock.Unlock()
	return c.count
}

/* ############################## ##
## #### Playground Functions #### ##
## ############################## */

func AtomicIntegerPlayground() {
	fmt.Println("Initializing counter := AtomicInteger{}")
	counter := AtomicInteger{}
	fmt.Println()

	fmt.Printf("counter.GetAndSet(1) = %d", counter.GetAndSet(1))
	fmt.Println()
	fmt.Printf("counter.Get() = %d", counter.Get())
	fmt.Println()

	fmt.Printf("counter.GetAndAdd(2) = %d", counter.GetAndAdd(2))
	fmt.Println()
	fmt.Printf("counter.Get() = %d", counter.Get())
	fmt.Println()

	fmt.Printf("counter.GetAndDecrement() = %d", counter.GetAndDecrement())
	fmt.Println()
	fmt.Printf("counter.Get() = %d", counter.Get())
	fmt.Println()

	fmt.Printf("counter.GetAndIncrement() = %d", counter.GetAndIncrement())
	fmt.Println()
	fmt.Printf("counter.Get() = %d", counter.Get())
	fmt.Println()

	doubler := func(v int) int { return (v * 2) }
	fmt.Printf("counter.GetAndUpdate(func(v int) int { return (v * 2) }) = %d", counter.GetAndUpdate(doubler))
	fmt.Println()
	fmt.Printf("counter.Get() = %d", counter.Get())
	fmt.Println()

	fmt.Printf("counter.AddAndGet(4) = %d", counter.AddAndGet(4))
	fmt.Println()

	fmt.Printf("counter.DecrementAndGet() = %d", counter.DecrementAndGet())
	fmt.Println()

	fmt.Printf("counter.IncrementAndGet() = %d", counter.IncrementAndGet())
	fmt.Println()

	fmt.Printf("counter.UpdateAndGet(func(v int) int { return (v * 2) }) = %d", counter.UpdateAndGet(doubler))
	fmt.Println()

	time.Sleep(time.Second)
}

func RunTwoGoroutinesIncrementingTheSameAtomicInteger() {
	fmt.Println("Initializing counter := AtomicInteger{}")
	counter := AtomicInteger{}
	fmt.Println()

	fmt.Println("Launch a goroutine named 'First Goroutine' that increment the counter a total of 5 times every 500 ms")
	go incrementBy(&counter, 5, 500, "First Goroutine")
	fmt.Println()

	fmt.Println("Launch a goroutine named 'Second Goroutine' that increment the counter a total of 5 times every 750 ms")
	go incrementBy(&counter, 5, 750, "Second Goroutine")
	fmt.Println()

	fmt.Println("Sleep for 4 second while the work gets done")
	time.Sleep(4 * time.Second)

	fmt.Printf("The expected final count is (expect 10): %d", counter.Get())
	fmt.Println()

	time.Sleep(time.Second)
}

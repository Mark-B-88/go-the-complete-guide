# Concurrency & Goroutines

First to understand what this is, we have to understand what it's not. Let's say that we have a 4 function calls in a normal program:

`store("hello")` - `store("hello")` - `store("hello")` - `sum(2,3)`

These functions are simply executed after each other. The 2nd function will only start after the 1st function finishes and so on and so on.

---

# Blocking the execution

Calling `slowGreet()` will block the execution of the last `greet()` call.

```go
package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello", phrase)
}

func main() {
	greet("Nice to meet you!")
	greet("How are you?")
	slowGreet("How...are...you...?")
	greet("Do you like the course?")
}
```

You can run code in parallel and functions concurrently by using a feature called **_Goroutines_**

---

# Running functions as goroutines

Adding `go` in front of a function will make it run in parallel. Notice though that when you run `go run .`, it finshes very fast, but does not console log anything?

```go
func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How...are...you...?")
	go greet("Do you like the course?")
}
```

That is because goroutines are lightweight threads managed by the Go runtine, and they execute independently of each other. However, the `main()` function exists as soon as it reaches the end, and when `main()` exits, the program terminates regardless of whether other goroutines have finished executing or not.

To allow the goroutines to complete their execution before the program exits, you can use synchronization mechanisms like **_channels_** or **_wait groups_**.

Here's an example:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(phrase string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second) // Simulate a slow, long-taking task
	fmt.Println("Hello", phrase)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4) // Number of goroutines to wait for

	go greet("Nice to meet you!", &wg)
	go greet("How are you?", &wg)
	go slowGreet("How...are...you...?", &wg)
	go greet("Do you like the course?", &wg)

	wg.Wait() // Wait for all goroutines to finish
}
```

---

# Using Channels

We can pass a **_channel_** as another parameter in our functions by providing it a name, then the `chan` keyword so that Go understands it's a **_channel_** and then the channel type, **_string, float, bool, etc_**

The `<-` operator is used for communication operations on channels. Its behavior depends on whether it's used on the left side or the right side of an expression.

`doneChan <- true` in this case denotes a **_send operation_**, this operation sends the value into the channel.

When `<-` is used to the right of the channel variable `(variable := <-channel)`, it denotes a **_receive operation_**.

```go
package main

import (
	"fmt"
	"time"
)

// func greet(phrase string) {
// 	fmt.Println("Hello!", phrase)
// }

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // Simulate a slow, long-taking task
	fmt.Println("Hello", phrase)
	doneChan <- true
}

func main() {
	// go greet("Nice to meet you!")
	// go greet("How are you?")
	done := make(chan bool)

	go slowGreet("How...are...you...?", done)
	// go greet("Do you like the course?")
	<- done
}
```

The channel `done` is used to coordinate the termination of the `slowGreet()` goroutine with the main goroutine, ensuring that the main goroutine waits until the slowGreet goroutine completes its task before proceeding further.

- `doneChan <- true`: This is a send operation, sending the boolean value true into the doneChan channel.

- `<-done`: This is a receive operation, receiving a value from the done channel. It waits until a value is sent into the channel.

---

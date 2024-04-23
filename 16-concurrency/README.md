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

Adding `go` in front of a function will make it run in parallel. Notice though that when you run `go run .`, it finshes very fast, but does not console log anything.

```go
func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How...are...you...?")
	go greet("Do you like the course?")
}
```

---

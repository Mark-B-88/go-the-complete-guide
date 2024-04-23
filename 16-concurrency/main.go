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

	wg.Add(4)

	go greet("Nice to meet you!", &wg)
	go greet("How are you?", &wg)
	go slowGreet("How...are...you...?", &wg)
	go greet("Do you like the course?", &wg)

	wg.Wait()
}

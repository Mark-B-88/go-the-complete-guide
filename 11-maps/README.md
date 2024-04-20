# How to Pretty Print a JSON Object

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	// Marshal the map to JSON with indentation
	prettyJSON, err := json.MarshalIndent(websites, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	// Print the formatted JSON
	fmt.Println(string(prettyJSON))
}
```

---

# Mutating Maps

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https:aws.com",
	}

	fmt.Println("Search by value: ", websites["Amazon Web Services"])

	// append another key / value to the map
	websites["LinkedIn"] = "https://linkedin.com"

	// deleting a value from the map
	delete(websites, "Google")


	// Pretty Print JSON
	prettyJSON, err := json.MarshalIndent(websites, "", "  ") // two tab indention
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}

```

---

# Maps vs Structs

Maps are way more flexible than structs and allow you to set labels to each value in the map. Structs are useful for pre-defined data such as a **_user_**, **_product_**, **_company_department_**, etc, etc.

---

# Using the special "make" function

Here's why you might want to use the make function in this context:

- Memory Allocation Efficiency:

When you use make to create a slice with an **_initial capacity_**, you are **_pre-allocating_** memory for that slice. This can help **_improve performance_**, especially if you know in advance the approximate size of the slice you'll need. By specifying the initial capacity, you avoid unnecessary **_memory allocations_** and resizing operations when appending elements to the slice later on. This can lead to better **_memory usage_** and reduced overhead.

- Avoiding Slice Resizing:

When you append elements to a slice beyond its initial length, Go automatically increases the slice's capacity as needed to accommodate the new elements. This involves allocating a new underlying array and copying existing elements, which can be inefficient if done frequently. By **_pre-allocating_** a larger capacity using `make()`, you can minimize the number of times Go needs to resize the slice's underlying array, improving performance.

- Control Over Memory Usage:

Using `make()` allows you to specify both the **_initial length_** and **_capacity_** of the slice, giving you more control over **_memory usage_**. You can **_allocate memory_** based on your application's specific requirements and optimize memory utilization accordingly.

In terms of use cases, you might want to use the make function when:

- You know in advance the approximate size of the slice you need, and you want to avoid frequent resizing operations.
- You want to optimize memory usage and reduce overhead by pre-allocating memory for a slice.
- You want more control over memory allocation and want to specify both the initial length and capacity of the slice.

The example below is not the correct **_use-case_** for the `make()` function.

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	userNames := make([]string, 2,5)

	userNames = append(userNames, "Mark")
	userNames = append(userNames, "Frank")

	prettyJSON, err := json.MarshalIndent(userNames, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
```

In a real life scenario, let's say you're developing a web application that handles user comments.

- Each comment consists of a text body, timestamp and the username of the commenter
- You expect the application to receive a large amount of comments
- You want to optimize memory useage and performance

Like this for example:

```go
package main

import (
    "fmt"
    "time"
)

// Comment represents a user comment.
type Comment struct {
    Text      string
    Timestamp time.Time
    Username  string
}

func main() {
    // Pre-allocate memory for storing comments
    const initialCapacity = 1000 // Estimated initial capacity
    comments := make([]Comment, 0, initialCapacity)

    // Simulate adding comments to the slice
    addComment(&comments, "Great post!", "user1")
    addComment(&comments, "Interesting topic!", "user2")
    addComment(&comments, "Thanks for sharing!", "user3")

    // Print the comments
    fmt.Println("Comments:")
    for _, comment := range comments {
        fmt.Printf("[%s] %s (%s)\n", comment.Timestamp.Format("2006-01-02 15:04:05"), comment.Text, comment.Username)
    }
}

// addComment adds a new comment to the comments slice.
func addComment(comments *[]Comment, text, username string) {
    // Create a new Comment instance
    comment := Comment{
        Text:      text,
        Timestamp: time.Now(),
        Username:  username,
    }

    // Append the new comment to the comments slice
    *comments = append(*comments, comment)
}

```

---

# Allocating Memory in Advance

If we know in advance how much space or slots we need in our map, we can simply use `make()` to pre-allocate the space.

If we keep 3 items, Go won't have to re-allocate memory. If we pass more than 3 or the limit assigned, then Go will have to re-allocate memory.

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	courseRatings := make(map[string]float64, 3)
	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.5 // map limit before re-allocating memory
	courseRatings["Node"] = 4.3 // re-allocate memory

	prettyJSON, err := json.MarshalIndent(courseRatings, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
```

# Using type aliases

You could also use a **_type alias_** to:

- Re-structure your code
- Export / Import functions
- Add a method to the type alias
- Make the code cleaner

```go
package main

import (
	"encoding/json"
	"fmt"
)

// type alias
type floatMap map[string]float64

func (m floatMap) Output() {
	prettyJSON, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}

func main() {
	courseRatings := floatMap{
		"Go":      4.7,
		"React":   4.8,
		"Angular": 4.5,
		"Node":    4.3,
	}

	courseRatings.Output()
}
```

---

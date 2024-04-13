# Understanding Pointers

This example results in too many copies in memory:

```go
package main

import "fmt"

func main() {
	age := 32 // regular variable
	fmt.Println("Age: ", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
```

This isn't a big deal since it's a small example, and even so the garbage collector will make sure to remove the copy when necessary. However, in a more complex function that requires a lot of calculations, this might become a problem to reply on copies instead of pointers.

---

# Creating a Pointer

```go
package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // this is the pointer
	agePointer = &age
	fmt.Println("Age: ", *agePointer) // this is the value of the pointer
}

func getAdultYears(age int) int {
	return age - 18
}
```

## How to assign the Pointer

You don't have to invoke the pointer like the example above, you can just do it like this:

```go
agePointer := &age
```

Intellisense in your IDE will pick up on the fact that it's a pointer, it will label it as `*int`

## How to get the value of the pointer

Adding an `*` in front of the pointer will reveal it's original value. If you omit the `*` it will reveal the address instead of the value.

```go
fmt.Println("Age: ", *agePointer) // original value
fmt.Println("Age: ", agePointer) // address in memory
```

## A Pointer's Null Value

All values in Go have a so-called "Null Value" - i.e., the value that's set as a default if no value is assigned to a variable.

For example, the null value of an int variable is 0. Of a float64, it would be 0.0. Of a string, it's "".

For a pointer, it's `nil` - a special value built-into Go.

`nil` represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.

---

# Using Pointers & Passing Them To Functions

```go
package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age
	fmt.Println("Age: ", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
```

---

# Using Pointers for Data Mutation

```go
package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age
	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18 // this will mutate the original amount of 32, to 14
}
```

---

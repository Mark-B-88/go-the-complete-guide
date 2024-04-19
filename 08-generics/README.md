# Interfaces, Dynamic Types & Limitations

Here is a small example of how you can use the "any" type to return almost any type. There are limitations and will be discussed in the next chapter. Also, this code is verbose and could be shorter with the use of **_generics_**

```go
package main

import "fmt"

func main() {
	result := add(1,2)
	result + 1
	// Go won't support this value because it doesn't understand what to do
}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aStr, aIsStr := a.(string)
	bStr, bIsStr := b.(string)

	if aIsStr && bIsStr {
		return aStr + bStr
	}
	return nil
}
```

The problem with this is that if we try to work with the results, Go won't understand how to return the results since it's values are "any", and they are too broad, for instance :

```go
func main() {
	result := add(1,2)

	result + 1
}
```

Go doesn't understand this value, and it won't understand what to do with the results. We can fix this with **_generics_**

# Introducing Generics

```go
package main

import "fmt"

func main() {
	// As an int
	result := add(1,2)
	result += 1

	// As a String
	result := add("Hello ", "World ")

	// As a float
	result := add(54.5, 63.9)

	fmt.Println(result)
}

func add[T int|float64|string](a, b T) T {
	return a + b
}
```

To turn the function into a **_generic function_**, you can do so by adding square brackets after the function name, and within those brackets, you can use a **_generic placeholder_**. Next to the placeholder, you can define which values it should accept and separate them with a pipe, and then also return the placeholder as well.

```go
func add[generic-type int|float|string](a, b generic-type) generic-type {
	return logic
}
```

You can also add more than one placeholder and return type if needed. This is for **_special use cases_** only and seldom should it be considered for use. If you don't use this wisely, you could potentially cause more issues further down the line.

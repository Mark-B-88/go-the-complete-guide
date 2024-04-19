# Understanding Struct Tags

If you add **_struct tags_**, in this case to json format, you can define custom titles for the JSON object.

```go
type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
```

This will result in the json format:

```json
{
	"title": "Lorem ipsum",
	"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In neque lorem, tristique eu urna id, suscipit vehicula augue. Sed eleifend aliquet scelerisque. Aenean id ex sed neque porttitor ornare. In et lobortis sem. Phasellus faucibus fermentum risus id efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce congue imperdiet ipsum, vel semper eros volutpat eget. In vitae velit sed mi mollis pulvinar. Quisque magna diam, auctor sit amet tristique ac, tristique vitae ligula. Nunc pharetra scelerisque auctor.",
	"created_at": "2024-04-15T21:12:16.8484352-06:00"
}
```

If you don't apply a **_struct tag_**, it will result in the json fields applying the same way as typed in the Struct itself.

---

# Creating an interface

Asides from creating an interface, you can also create an **_embedded_** interface, which can take the characteristics of other interfaces, as well as define new methods.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// embedded interface
type outputtable interface {
	saver
	Display()
	// DoSomething(int) string - method(value) - return value
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}

	fmt.Println("saving the note succeeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
```

---

# The Special "Any Value Allowed" Type

```go
func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello World")
}

// The Special "Any Value Allowed" type
// any value is allowed, int, string, etc, etc
// you can either state interface {} or "any"
func printSomething(value interface{}) {
	fmt.Println(value)
}
```

---

# Extrating Type Information From Values

This `value.(value-type)` can help you extract type information from values. The naming convention is not mandatory, you can name it depending on your **_use-case_**.

And that's what this means:

`valueType, boolean-value := value.(value-type)`
`intVal, ok := value.(int)`

You can make conditionals based on either the true or false response.

```go
func printSomething(value interface{}) {
	// valueType, boolean-value := value.(value-type)
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float: ", floatVal)
		return
	}

	strVal, ok := value.(string)

	if ok {
		fmt.Println(strVal)
		return
	}
}
```

You could also do a switch statement style:

```go
func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println(value)
	}
}
```

You could also set a default, or not, it depends on your **_use-case_**.

---

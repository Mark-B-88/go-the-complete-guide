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
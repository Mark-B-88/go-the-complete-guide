package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// userNames := make([]string, 2,5)

	// userNames = append(userNames, "Mark")
	// userNames = append(userNames, "Frank")

	courseRatings := make(map[string]float64, 3)
	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.5
	courseRatings["Node"] = 4.3

	prettyJSON, err := json.MarshalIndent(courseRatings, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
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
	// userNames := make([]string, 2,5)

	// userNames = append(userNames, "Mark")
	// userNames = append(userNames, "Frank")

	// courseRatings := make(floatMap, 3)
	// courseRatings["Go"] = 4.7
	// courseRatings["React"] = 4.8
	// courseRatings["Angular"] = 4.5
	// courseRatings["Node"] = 4.3

	courseRatings := floatMap{
		"Go":      4.7,
		"React":   4.8,
		"Angular": 4.5,
		"Node":    4.3,
	}

	courseRatings.Output()
}
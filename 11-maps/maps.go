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

package main

import "fmt"

type Product struct {
	id string
	title string
	price float64
}

func main() {
	// create an array that contains 3 hobbies
	hobbies := [3]string{"music", "cooking", "reading"}
	fmt.Println(hobbies)

	// output more data about the array
	fmt.Println(hobbies[0]) // 1st element - music
	fmt.Println(hobbies[1:3]) // 2nd & 3rd element - cooking, reading - more explicit
	// fmt.Println(hobbies[1:]) // 2nd & 3rd element - cooking, reading - rest of array

	// another variable based on a slice
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// displays the amount of items in the array
	fmt.Println(cap(mainHobbies)) // 3

	// 2nd and last element
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// create a dynamic array
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// re-assign the 2nd item
	courseGoals[1] = "Learn all the details!"

	// assgin an new item
	courseGoals = append(courseGoals, "Build a REST API")

	fmt.Println(courseGoals)

	// dynamic list of products with struct
	products := []Product{
		{
			"d2b07d37-82d0-439c-b0cd-9f6b2ad56d70",
			"Generic 40gb USB flash drive",
			29.99,
		},
		{
			"f0218128-b341-44d0-ba94-4809f9e57e5c",
			"Samsung 2tb SSD",
			149.99,
		},
	}
	fmt.Println(products)

	// add a new product
	newProduct := Product{
		"2a057ec0-8546-4904-8fd8-b1bdf56f3e82",
		"AMD-Ryzen 5 CPU",
		299.99,
	}

	// append new product
	products = append(products, newProduct)
	fmt.Println(products)
}
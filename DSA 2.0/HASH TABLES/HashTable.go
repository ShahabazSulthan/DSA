package main

import "fmt"

func hashTable() {
	// Create a map with string keys and int values
	age := make(map[string]int)

	// Add key-value pairs to the map
	age["Alice"] = 25
	age["Bob"] = 30
	age["Charlie"] = 35

	// Access a value by its key
	fmt.Println("Alice's age is", age["Alice"])

	// Check if a key exists
	if value, exists := age["Bob"]; exists {
		fmt.Println("Bob's age is", value)
	} else {
		fmt.Println("Bob's age not found")
	}

	// Delete a key-value pair
	delete(age, "Charlie")

	// Iterate over the map
	for key, value := range age {
		fmt.Printf("%s is %d years old\n", key, value)
	}
}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func structMap() {
	// Create a map with string keys and Person values
	people := make(map[string]Person)

	// Add key-value pairs to the map
	people["alice"] = Person{Name: "Alice", Age: 25}
	people["bob"] = Person{Name: "Bob", Age: 30}
	people["charlie"] = Person{Name: "Charlie", Age: 35}

	// Update a value in the map
	people["bob"] = Person{Name: "Bob", Age: 31}

	// Check if a key exists and print its value
	if person, exists := people["bob"]; exists {
		fmt.Println("Bob's updated age is", person.Age)
	} else {
		fmt.Println("Bob not found")
	}

	// Attempt to access a non-existing key
	if person, exists := people["dave"]; exists {
		fmt.Println("Dave's age is", person.Age)
	} else {
		fmt.Println("Dave not found")
	}

	// Delete a key-value pair
	delete(people, "charlie")

	// Iterate over the map and print all entries
	for key, person := range people {
		fmt.Printf("%s: %s is %d years old\n", key, person.Name, person.Age)
	}
}

package main

import "fmt"

// Person struct defines the properties of a person
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Address struct to demonstrate nested structs
type Address struct {
	Street  string
	City    string
	Country string
}

// Employee demonstrates nested structs
type Employee struct {
	Details Person
	WorkLoc Address
}

func main() {
	fmt.Println("--- Struct Basics ---")

	// 1. Initializing with field names
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john@example.com",
	}
	fmt.Printf("Person 1: %+v\n", p1)

	// 2. Initializing without field names (not recommended for large structs)
	p2 := Person{"Jane", "Smith", 25, "jane@example.com"}
	fmt.Println("Person 2:", p2.FirstName, p2.LastName)

	// 3. Zero values
	var p3 Person
	fmt.Printf("Zero value Person: %+v\n", p3)
	p3.FirstName = "Alice"
	fmt.Println("Updated p3:", p3.FirstName)

	// 4. Pointer to a struct
	p4 := &Person{FirstName: "Bob", Age: 40}
	fmt.Println("Pointer to Person:", p4.FirstName) // No need to dereference like (*p4).FirstName

	// 5. Anonymous structs
	config := struct {
		APIKey string
		Port   int
	}{
		APIKey: "secret123",
		Port:   8080,
	}
	fmt.Printf("Config: %+v\n", config)

	// 6. Nested structs
	emp := Employee{
		Details: Person{FirstName: "Sam", LastName: "Altman"},
		WorkLoc: Address{Street: "123 Tech Ln", City: "San Francisco", Country: "USA"},
	}
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Println("Employee City:", emp.WorkLoc.City)
}

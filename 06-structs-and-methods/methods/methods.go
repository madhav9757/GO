package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

// 1. Value Receiver Method
// It operates on a copy of the Circle. Original value won't change.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 2. Pointer Receiver Method
// It operates on the actual memory address. Can modify the original value.
func (c *Circle) Scale(factor float64) {
	c.Radius = c.Radius * factor
}

// A more complex example with a Person struct
type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) Display() string {
	return fmt.Sprintf("[%d] %s <%s>", u.ID, u.Name, u.Email)
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	fmt.Println("--- Methods in Go ---")

	c := Circle{Radius: 5}
	fmt.Printf("Initial Circle: %+v, Area: %.2f\n", c, c.Area())

	// Calling a pointer receiver method
	c.Scale(2)
	fmt.Printf("After Scaling: %+v, Area: %.2f\n", c, c.Area())

	fmt.Println("\n--- User Example ---")
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	fmt.Println("Original:", user.Display())

	user.UpdateEmail("alice.new@example.com")
	fmt.Println("Updated :", user.Display())
}

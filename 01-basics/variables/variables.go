package main

import "fmt"

func main() {
	// 1. Variable declaration with type
	var a int = 10
	var name string = "Madhav"

	// 2. Variable declaration without type (type inference)
	var b = 20
	var city = "Delhi"

	// 3. Short-hand declaration (inside function only)
	c := 30
	isGo := true

	// 4. Multiple variable declaration
	var x, y, z int = 1, 2, 3

	// 5. Multiple short-hand declaration
	p, q := 5, "Go"

	// 6. Zero value variables
	var n int
	var s string
	var f float64
	var flag bool

	fmt.Println("a =", a)
	fmt.Println("name =", name)
	fmt.Println("b =", b)
	fmt.Println("city =", city)
	fmt.Println("c =", c)
	fmt.Println("isGo =", isGo)
	fmt.Println("x, y, z =", x, y, z)
	fmt.Println("p, q =", p, q)

	fmt.Println("\nZero values:")
	fmt.Println("int:", n)
	fmt.Println("string:", s)
	fmt.Println("float:", f)
	fmt.Println("bool:", flag)
}

package main

import "fmt"

func types() {
	// Integer types
	var a int = 10
	var b int8 = 100
	var c uint = 20

	// Floating point types
	var x float32 = 3.14
	var y float64 = 9.81

	// Boolean type
	var isGoFun bool = true

	// String type
	var name string = "Golang"

	// Character (rune)
	var ch rune = 'A' // rune is alias for int32

	// Byte
	var bt byte = 255 // byte is alias for uint8

	fmt.Println("Integer:", a, b, c)
	fmt.Println("Float:", x, y)
	fmt.Println("Boolean:", isGoFun)
	fmt.Println("String:", name)
	fmt.Println("Rune (char):", ch)
	fmt.Println("Byte:", bt)
}

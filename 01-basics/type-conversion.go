package main

import "fmt"

func type_conversion() {
	var a int = 10
	var b float64 = 3.5

	// Type Conversion
	var sum float64 = float64(a) + b

	fmt.Println("Integer value:", a)
	fmt.Println("Float value:", b)
	fmt.Println("Sum after type conversion:", sum)

	// int to float
	var x int = 25
	var y float64 = float64(x)
	fmt.Println("Int to Float:", y)

	// float to int (decimal part lost)
	var p float64 = 9.8
	var q int = int(p)
	fmt.Println("Float to Int:", q)
}

package main

import (
	"fmt"

	modulea "github.com/madhav9757/module-a"
)

func main() {
	fmt.Println(modulea.Greet("Go Programmer"))
	fmt.Println("\nThis project demonstrates how to use local modules.")
	fmt.Println("Check the go.mod file in this directory to see the 'replace' directive.")
}

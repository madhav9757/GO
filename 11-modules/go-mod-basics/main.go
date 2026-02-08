package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// Print standard text
	fmt.Println("This is standard white text.")

	// Use the color package to print colored text
	color.Cyan("This is cyan text using the 'color' module!")
	color.Green("This is green text, which means everything is working fine.")

	// You can also mix colors and styles
	c := color.New(color.FgBlue).Add(color.Underline)
	c.Println("This is underlined blue text.")

	fmt.Println("\nTo use this module:")
	fmt.Println("1. Run 'go mod init <module-name>' (already done)")
	fmt.Println("2. Run 'go get github.com/fatih/color' (already done)")
	fmt.Println("3. Run 'go run main.go'")
}

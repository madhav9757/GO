package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "example.txt"

	// 1. Write the whole file at once
	fmt.Println("Writing to file...")
	message := "Line 1: Hello Go!\nLine 2: File handling is easy.\n"
	err := os.WriteFile(fileName, []byte(message), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	// 2. Open file for appending
	fmt.Println("Appending to file...")
	// os.O_APPEND: append data to the file when writing.
	// os.O_WRONLY: open the file write-only.
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening for append: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("Line 3: This line was appended.\n"); err != nil {
		fmt.Printf("Error appending string: %v\n", err)
	}

	fmt.Println("File operations completed successfully.")
}

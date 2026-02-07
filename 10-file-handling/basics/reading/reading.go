package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. Reading an entire file into memory (Best for small files)
	fmt.Println("--- Reading entire file ---")
	content, err := os.ReadFile("example.txt")
	if err != nil {
		// If the file doesn't exist, we'll create it for the next run
		fmt.Printf("Error reading file: %v\n", err)
	} else {
		fmt.Printf("File Content: %s\n", string(content))
	}

	// 2. Reading line by line using a scanner (Best for large files/logs)
	fmt.Println("\n--- Reading line by line ---")
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	} else {
		defer file.Close() // Ensure the file is closed!

		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
			lineNum++
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error during scanning: %v\n", err)
		}
	}

	// 3. Reading in chunks with a buffer
	fmt.Println("\n--- Reading in chunks ---")
	f, err := os.Open("example.txt")
	if err == nil {
		defer f.Close()
		buffer := make([]byte, 16) // 16 byte buffer
		for {
			n, err := f.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Error reading: %v\n", err)
				break
			}
			fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
		}
	}
}

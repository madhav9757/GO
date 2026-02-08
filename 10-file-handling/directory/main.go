package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dirName := "test_dir"

	// 1. Creating a directory
	fmt.Printf("Creating directory: %s\n", dirName)
	err := os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error: %v\n", err)
	}

	// 2. Creating nested directories
	nestedDir := filepath.Join(dirName, "level1", "level2")
	fmt.Printf("Creating nested directories: %s\n", nestedDir)
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 3. Creating a dummy file inside the nested directory
	filePath := filepath.Join(nestedDir, "temp.txt")
	os.WriteFile(filePath, []byte("Self-destructing in 3... 2... 1..."), 0644)

	// 4. Walking the directory tree
	fmt.Println("\n--- Walking the directory structure ---")
	err = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("Found: %s (Is Directory: %v)\n", path, info.IsDir())
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking: %v\n", err)
	}

	// 5. Cleanup: Removing everything we created
	fmt.Printf("\nCleaning up: Deleting %s\n", dirName)
	// Be careful with RemoveAll!
	err = os.RemoveAll(dirName)
	if err != nil {
		fmt.Printf("Error during cleanup: %v\n", err)
	}
}

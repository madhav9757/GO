# 📁 File Handling in Go

File handling is a fundamental aspect of many applications, from logging data to processing configuration files or persistent storage. Go provides a powerful and efficient set of packages for interacting with the file system.

## 📌 Core Concepts

### 1. The `os` Package

The `os` package provides a platform-independent interface to operating system functionality. This is the primary package used for opening, creating, and managing files.

- `os.Create(name string)`: Creates a new file or truncates an existing one.
- `os.Open(name string)`: Opens a file for reading.
- `os.OpenFile(name, flag, perm)`: Opens a file with specific flags (e.g., append mode) and permissions.

### 2. Reading and Writing

There are several ways to read and write files in Go:

- **Quick Read/Write**: Using `os.ReadFile` and `os.WriteFile` (formerly in `io/ioutil`) is the simplest way to handle entire files at once.
- **Buffered IO**: The `bufio` package is used for efficient reading and writing by reducing the number of system calls.
- **Streaming**: For large files, using `io.Copy` or manual buffers with `file.Read` is more memory-efficient.

### 3. JSON Handling

Go has built-in support for JSON through the `encoding/json` package.

- **Marshaling**: Converting Go structs/maps to JSON strings (`json.Marshal`).
- **Unmarshaling**: Converting JSON strings back into Go types (`json.Unmarshal`).

---

## 📂 Module Structure

- **[basics](./basics)**: Introduction to reading and writing text files.
- **[json](./json)**: Working with JSON data structures.
- **[directory](./directory)**: Creating directories and managing file paths.

---

## 🚀 Getting Started

### Creating and Writing to a File

```go
package main

import (
    "os"
)

func main() {
    content := []byte("Hello, Go File Handling!")
    err := os.WriteFile("test.txt", content, 0644)
    if err != nil {
        panic(err)
    }
}
```

### Reading a File

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data))
}
```

---

## 💡 Best Practices

1. **Always Close Files**: Use `defer file.Close()` immediately after successfully opening a file to ensure resources are freed.
2. **Handle Errors**: File operations are prone to errors (permissions, missing files, disk full). Always check the returned `err`.
3. **Use Buffered IO for Large Files**: Avoid loading massive files into memory at once. Use `bufio.Scanner` or `bufio.Reader`.
4. **File Permissions**: Use appropriate permissions (e.g., `0644` for files, `0755` for directories).

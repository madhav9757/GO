# Hello World in Go

This is a simple **Hello World** program written in **Go (Golang)**.  
It demonstrates the basic structure of a Go program and how to print output to the console.

---

## 📂 File

- `hello.go` – Prints "Hello, World!" to the console

---

## 🧾 Program Code

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## ⚙️ Explanation

- `package main`
  Defines the main package. Every executable Go program must have this.

- `import "fmt"`
  Imports the fmt package, which provides formatted I/O functions.

- `func main()`
  Entry point of the Go program. Execution starts here.

- `fmt.Println()`
  Prints text to the console followed by a new line.

---

## How to Run

1. Make sure Go is installed on your system.

2. Open terminal in the folder containing hello.go.

3. Run the program:

```bash
go run hello.go
```

---

## Output

```
Hello, World!
```


---

## 
📌 Notes
- This is usually the first program written while learning Go.
- Each Go executable program must contain exactly one main() function.
---



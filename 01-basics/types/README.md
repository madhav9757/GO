# Go Data Types Program

This Go program demonstrates the **basic data types in Go**, including integers, floats, booleans, strings, runes (characters), and bytes.

---

## 📂 File

- `types.go` – Main program file demonstrating Go data types

---

## 🧾 Program Code

```go
package main

import "fmt"

func main() {
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
```

---

## ⚙️ Explanation

**Integer Types**
- `int` – Platform-dependent size (32 or 64 bits)
- `int8` – 8-bit signed integer (-128 to 127)
- `int16` – 16-bit signed integer
- `int32` – 32-bit signed integer
- `int64` – 64-bit signed integer
- `uint` – Unsigned platform-dependent integer
- `uint8`, `uint16`, `uint32`, `uint64` – Unsigned variants

**Floating Point Types**
- `float32` – 32-bit floating point number
- `float64` – 64-bit floating point number (more precise)

**Boolean Type**
- `bool` – True or false values only

**String Type**
- `string` – Sequence of characters (UTF-8 encoded)

**Special Types**
- `rune` – Alias for `int32`, represents a Unicode code point
- `byte` – Alias for `uint8`, represents a single byte of data

---

## ⚡ How to Run

1. Make sure Go is installed on your system.
2. Open terminal and navigate to the folder containing `types.go`.
3. Run the program:
   ```bash
   go run types.go
   ```

---

## ▶️ Expected Output

```
Integer: 10 100 20
Float: 3.14 9.81
Boolean: true
String: Golang
Rune (char): 65
Byte: 255
```

---

## 🔑 Key Points

- Go is a **statically typed** language – types must be known at compile time
- `int` and `uint` sizes depend on the system (32-bit or 64-bit)
- `float64` is more precise than `float32` and is preferred for most calculations
- `rune` is used for Unicode characters and prints as integer (ASCII/Unicode value)
- `byte` is commonly used for raw data and binary operations
- Strings in Go are **immutable** and UTF-8 encoded by default
- Character 'A' has ASCII/Unicode value of 65

---

## 📊 Type Size Reference

| Type | Size | Range |
|------|------|-------|
| int8 | 1 byte | -128 to 127 |
| uint8 (byte) | 1 byte | 0 to 255 |
| int16 | 2 bytes | -32,768 to 32,767 |
| uint16 | 2 bytes | 0 to 65,535 |
| int32 (rune) | 4 bytes | -2,147,483,648 to 2,147,483,647 |
| uint32 | 4 bytes | 0 to 4,294,967,295 |
| int64 | 8 bytes | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| uint64 | 8 bytes | 0 to 18,446,744,073,709,551,615 |
| float32 | 4 bytes | ±3.4e±38 |
| float64 | 8 bytes | ±1.7e±308 |

---

## 💡 Best Practices

- Use `int` for general integer operations unless you need a specific size
- Use `float64` for floating-point calculations (default for decimal numbers)
- Use `rune` when working with Unicode characters
- Use `byte` when working with ASCII or raw binary data
- Choose the smallest type that fits your needs to save memory

---
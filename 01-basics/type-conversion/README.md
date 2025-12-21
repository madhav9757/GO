# Go Type Conversion Program

This Go program demonstrates **type conversion in Go**. It shows how to convert between different numeric types (integers and floats) and explains the behavior of such conversions.

---

## 📂 File

- `type-conversion.go` – Main program file demonstrating type conversion in Go

---

## 🧾 Program Code

```go
package main

import "fmt"

func main() {
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
```

---

## ⚙️ Explanation

**package main**  
Required for executable Go programs.

**import "fmt"**  
Used for printing output to the console.

**Type Conversion Syntax**  
`targetType(value)` – Converts value to targetType

**int to float64**  
`float64(a)` – Converts integer to float, preserves value exactly

**float64 to int**  
`int(p)` – Converts float to integer, **truncates decimal part** (not rounded)

---

## ⚡ How to Run

1. Make sure Go is installed on your system.
2. Open terminal and navigate to the folder containing `type-conversion.go`.
3. Run the program:
   ```bash
   go run type-conversion.go
   ```

---

## ▶️ Expected Output

```
Integer value: 10
Float value: 3.5
Sum after type conversion: 13.5
Int to Float: 25
Float to Int: 9
```

---

## 🔑 Key Points

- **Go does not support implicit type conversion** – you must explicitly convert types
- Converting `int` to `float` preserves the value exactly
- Converting `float` to `int` **truncates** (removes) the decimal part – it does **not** round
- Type conversion uses the syntax: `type(value)`
- Mixed-type arithmetic requires explicit conversion
- Always ensure the target type can hold the value to avoid overflow

---

## ⚠️ Important Notes

- **9.8** converted to `int` becomes **9** (not 10)
- Converting large floats to smaller int types can cause overflow
- Go will throw a compile-time error if you try to mix types without conversion

---
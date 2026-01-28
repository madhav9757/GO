# Pointers in Go

Pointers are a fundamental concept in Go that allow you to work directly with memory addresses. Using pointers, you can pass references to values instead of copying the values themselves.

## 📂 Directory Structure

```
05-pointers/
├── basics/
│   ├── pointers.go
│   └── README.md
├── pointer-functions/
│   ├── pointer-functions.go
│   └── README.md
└── README.md
```

## 📚 Topics Covered

### 1. Basic Pointers (`basics/`)
Learn the fundamental syntax and concepts:
- Memory addresses and the address-of operator `&`
- Pointer types and the dereference operator `*`
- The `nil` value for pointers
- The `new()` function

### 2. Pointers in Functions (`pointer-functions/`)
Learn when and why to use pointers with functions:
- Pass by Value vs. Pass by Reference
- Modifying original variables from within functions
- Efficiency when dealing with large data structures

## 🚀 Getting Started

Each subdirectory contains:
- A `.go` file with example code
- A `README.md` with detailed explanations

To run any example:

```bash
# Navigate to a specific folder
cd basics/

# Run the example
go run pointers.go
```

## 🎯 Key Takeaways

1. **Memory Efficiency:** Pointers allow you to pass large structs or arrays without copying the entire data structure.
2. **Mutations:** Use pointers when you need a function to modify the original value of an argument.
3. **No Pointer Arithmetic:** Unlike C/C++, Go does not allow pointer arithmetic (e.g., `p++`) for safety reasons, unless you use the `unsafe` package.
4. **The `nil` Pointer:** Always check if a pointer is `nil` before dereferencing it to avoid runtime panics.

## 📖 Learn More

- [Go by Example: Pointers](https://gobyexample.com/pointers)
- [A Tour of Go: Pointers](https://go.dev/tour/moretypes/1)
- [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)

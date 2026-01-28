# Functions in Go

Functions are the building blocks of Go programs. They allow you to structure your code into reusable logic blocks.

## 📂 Directory Structure

```
04-functions/
├── basics/
│   ├── functions.go
│   └── README.md
├── multiple-returns/
│   ├── multiple-returns.go
│   └── README.md
├── variadic/
│   ├── variadic.go
│   └── README.md
├── anonymous/
│   ├── anonymous.go
│   └── README.md
├── closures/
│   ├── closures.go
│   └── README.md
└── README.md
```

## 📚 Topics Covered

### 1. Basic Functions (`basics/`)
Learn how to define and call functions:
- Function declaration syntax
- Parameters and return types
- Named return values

### 2. Multiple Returns (`multiple-returns/`)
Go supports returning multiple values from a function:
- Syntax for returning multiple values
- Handling multiple return values
- Ignoring values with the blank identifier `_`

### 3. Variadic Functions (`variadic/`)
Functions that accept a variable number of arguments:
- Usage of `...Type` syntax
- Passing slices to variadic functions
- Applications (e.g., `fmt.Println`)

### 4. Anonymous Functions (`anonymous/`)
Functions without a name (lambdas):
- Defining inline functions
- Assigning functions to variables
- Immediate execution

### 5. Closures (`closures/`)
Functions that capture their surrounding state:
- Understanding variable scope and lifetime
- Creating factory functions
- State isolation

## 🚀 Getting Started

Each subdirectory contains:
- A `.go` file with example code
- A `README.md` with detailed explanations

To run any example:

```bash
# Navigate to a specific folder
cd basics/

# Run the example
go run functions.go
```

## 🎯 Key Takeaways

1. **First-Class Citizens:** Functions in Go are values. They can be passed as arguments, returned directly, and assigned to variables.
2. **Multiple Returns:** A powerful feature often used for returning a result and an error (e.g., `result, err := doSomething()`).
3. **Pass by Value:** Arguments are passed by value (copying data). Use pointers to modify the original data.
4. **Defer:** Use `defer` to ensure function calls (like cleanup) happen before the surrounding function returns.

## 📖 Learn More

- [Go by Example: Functions](https://gobyexample.com/functions)
- [Go by Example: Multiple Return Values](https://gobyexample.com/multiple-return-values)
- [Go by Example: Variadic Functions](https://gobyexample.com/variadic-functions)
- [Go by Example: Closures](https://gobyexample.com/closures)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)

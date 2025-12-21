# 01-basics - Go Programming Fundamentals

This folder contains **basic Go programming examples** covering fundamental concepts for beginners. Each subfolder focuses on a specific core concept with working code examples and detailed documentation.

---

## 📁 Folder Structure

```
01-basics/
├── hello-world/
│   ├── hello-world.go
│   └── README.md
├── variables/
│   ├── variables.go
│   └── README.md
├── types/
│   ├── types.go
│   └── README.md
├── type-conversion/
│   ├── type-conversion.go
│   └── README.md
└── input-output/
    ├── input-output.go
    └── README.md
```

---

## 📚 Topics Covered

### 1. **Hello World**
📂 `hello-world/`

Your first Go program! Learn the basic structure of a Go program and how to print output to the console.

**Key Concepts:**
- Package declaration (`package main`)
- Import statements
- The `main()` function
- Using `fmt.Println()`

---

### 2. **Variables**
📂 `variables/`

Learn different ways to declare and use variables in Go.

**Key Concepts:**
- Variable declaration with explicit types
- Type inference
- Short-hand declaration (`:=`)
- Multiple variable declarations
- Zero values for uninitialized variables

**Example:**
```go
var a int = 10          // Explicit type
var b = 20              // Type inference
c := 30                 // Short-hand
var x, y, z = 1, 2, 3  // Multiple variables
```

---

### 3. **Data Types**
📂 `types/`

Explore the fundamental data types available in Go.

**Key Concepts:**
- Integer types (`int`, `int8`, `int16`, `int32`, `int64`, `uint`)
- Floating point types (`float32`, `float64`)
- Boolean type (`bool`)
- String type (`string`)
- Special types (`rune`, `byte`)

**Example:**
```go
var age int = 25
var pi float64 = 3.14159
var isActive bool = true
var name string = "Go"
var ch rune = 'A'
```

---

### 4. **Type Conversion**
📂 `type-conversion/`

Learn how to convert between different data types in Go.

**Key Concepts:**
- Explicit type conversion (Go doesn't support implicit conversion)
- Converting between numeric types
- Understanding truncation when converting float to int
- Mixed-type arithmetic

**Example:**
```go
var a int = 10
var b float64 = float64(a)  // int to float
var c float64 = 9.8
var d int = int(c)          // float to int (9, not 10!)
```

---

### 5. **Input and Output**
📂 `input-output/`

Learn how to take user input from the keyboard and display output.

**Key Concepts:**
- Reading user input with `fmt.Scan()`
- Using the address operator (`&`)
- Formatted output with `fmt.Print()` and `fmt.Println()`
- Basic user interaction

**Example:**
```go
var name string
fmt.Print("Enter your name: ")
fmt.Scan(&name)
fmt.Println("Hello,", name)
```

---

## 🚀 Getting Started

### Prerequisites
- Go installed on your system (version 1.16 or higher recommended)
- Basic understanding of programming concepts
- A text editor or IDE (VS Code, GoLand, etc.)

### Running the Programs

Navigate to any subfolder and run the program:

```bash
cd hello-world
go run hello-world.go
```

Or run from the parent directory:

```bash
go run 01-basics/hello-world/hello-world.go
```

---

## 📖 Learning Path

Follow this recommended order for best learning experience:

1. **Start with Hello World** → Understand basic program structure
2. **Move to Variables** → Learn how to store and manage data
3. **Explore Data Types** → Understand different types of data
4. **Practice Type Conversion** → Learn to work with multiple types
5. **Try Input/Output** → Create interactive programs

---

## 💡 Key Takeaways

After completing this section, you should understand:

✅ Basic Go program structure and syntax  
✅ How to declare and use variables  
✅ Different data types and when to use them  
✅ How to convert between types safely  
✅ How to get input from users and display output  
✅ Go's zero values for different types  
✅ The importance of explicit type conversion  

---

## 🔑 Important Go Basics Rules

- Every executable Go program must have a `package main`
- Every executable program needs exactly one `main()` function
- Go is **strongly typed** – types must be known at compile time
- Go does **not** support implicit type conversion
- Unused variables cause **compile-time errors**
- Short-hand declaration (`:=`) can only be used **inside functions**
- Variables have **zero values** by default (0 for numbers, "" for strings, false for bools)

---

## 📝 Notes

- Each subfolder contains its own README with detailed explanations
- All programs are standalone and can be run independently
- Code examples are kept simple for educational purposes
- Comments are included in the code for clarity

---

## 🎯 Next Steps

After mastering these basics, you can move on to:
- Control structures (if/else, switch, loops)
- Functions and methods
- Arrays and slices
- Maps and structs
- Pointers
- Error handling

---

## 🤝 Contributing

Feel free to improve these examples or add new basic concepts!

---

**Happy Learning! 🎉**
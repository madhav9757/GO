# Structs and Methods in Go

Structs are Go's way of creating custom, complex types by grouping fields together. Methods add behavior to these types, enabling an object-oriented style of programming.

## 📂 Directory Structure

```
06-structs-and-methods/
├── basics/
│   ├── structs.go
│   └── README.md
├── methods/
│   ├── methods.go
│   └── README.md
├── embedding/
│   ├── embedding.go
│   └── README.md
└── README.md
```

## 📚 Topics Covered

### 1. Basic Structs (`basics/`)
Learn the fundamental syntax:
- Defining struct types
- Initialization methods (Field names vs Positional)
- Zero values and Pointers
- Anonymous and Nested structs

### 2. Methods (`methods/`)
Learn how to define behavior:
- Value receivers vs. Pointer receivers
- Mutating struct state
- Guidelines for choosing the right receiver

### 3. Composition via Embedding (`embedding/`)
Learn Go's approach to code reuse:
- Anonymous field embedding
- Field and Method promotion
- Shadowing (Overriding) behavior
- Composition vs Inheritance

## 🚀 Getting Started

Each subdirectory contains:
- A `.go` file with commented example code
- A `README.md` with detailed explanations

To run any example:

```bash
# Navigate to a specific folder
cd basics/

# Run the example
go run structs.go
```

## 🎯 Key Takeaways

1. **Clean Data Structures:** Structs help you organize related data cleanly.
2. **Behavioral Attachment:** Methods allow types to satisfy interfaces, which is the key to polymorphism in Go.
3. **Pointer Receivers:** Use pointer receivers if you need to modify the state or if the struct is large (to avoid copying).
4. **Prefer Composition:** Use embedding to build complex types from simpler ones rather than deep inheritance hierarchies.

## 📖 Learn More

- [Go by Example: Structs](https://gobyexample.com/structs)
- [Go by Example: Methods](https://gobyexample.com/methods)
- [Go by Example: Embedding](https://gobyexample.com/struct-embedding)
- [Effective Go: Methods](https://go.dev/doc/effective_go#methods)

# Struct Embedding (Composition)

Go does not have class-based inheritance. Instead, it promotes code reuse through **Composition** using **Struct Embedding**.

## 📝 Key Concepts

### 1. What is Embedding?
Embedding happens when you include a type within a struct without a field name.
```go
type Admin struct {
    User // Embedded field
    Role string
}
```

### 2. Field Promotion
Fields of the embedded struct are "promoted" to the outer struct. You can access them directly as if they were declared in the outer struct.
- `admin.Name` (if `User` has `Name`)
- `admin.User.Name` (Explicit access still works)

### 3. Method Promotion
Methods of the embedded type are also promoted.
- If `User` has a `Login()` method, `admin.Login()` works automatically.

### 4. Shadowing
If the outer struct defines a field or method with the same name as an embedded one, the outer one "shadows" the inner one.
- You can still access the inner one explicitly: `admin.User.MethodName()`.

### 5. Why not Inheritance?
Composition is often more flexible than inheritance. It creates a "has-a" or "includes-a" relationship rather than an "is-a" relationship, which leads to flatter and more maintainable code structures.

## 🚀 How to Run
```bash
go run embedding.go
```

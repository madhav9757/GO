# Methods in Go

Methods are functions attached to a specific type (usually a struct). They allow you to define behavior for your data structures.

## 📝 Key Concepts

### 1. Pointer vs Value Receivers

| Feature | Value Receiver `(t T)` | Pointer Receiver `(t *T)` |
|---------|------------------------|---------------------------|
| **Data** | Operates on a **copy** | Operates on the **original** |
| **Mutation**| Cannot modify fields | Can modify fields |
| **Performance**| Copy overhead for large structs | Memory efficient (shares address) |

### 2. Method Declaration
```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### 3. Why use Methods?
- **Encapsulation:** Group data and behavior together.
- **Interfaces:** Methods are essential for implementing interfaces (polymorphism).
- **Readability:** `circle.Area()` is more intuitive than `computeArea(circle)`.

### 4. Method Sets
- A type `T` can access methods defined with both `(t T)` and `(t *T)` because Go automatically handles the conversion.
- However, for interfaces, the method set counts. (More on this in the Interfaces module).

## 🚀 How to Run
```bash
go run methods.go
```

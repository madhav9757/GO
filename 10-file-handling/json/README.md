# JSON Handling in Go

Go provides excellent support for JSON via the `encoding/json` package. This section covers how to convert Go data structures to JSON (Marshaling) and back (Unmarshaling).

## 📖 Key Concepts

### 1. Marshaling

Converting a Go `struct` or `map` into a JSON string.

```go
data, _ := json.Marshal(user)
```

### 2. Unmarshaling

Parsing a JSON string into a Go `struct` or `map`.

```go
json.Unmarshal(data, &user)
```

### 3. Struct Tags

Customizing JSON keys using tags:

```go
type User struct {
    Name  string `json:"full_name"`
    Email string `json:"email_address,omitempty"`
}
```

## 🛠️ Files

- `main.go`: Demonstrates marshaling, unmarshaling, and working with complex structures.
- `users.json`: A sample data file used for unmarshaling examples.

# 📦 Go Modules and Dependency Management

Go Modules are the standard dependency management system in Go. They allow you to manage project dependencies, specify versions, and ensure reproducible builds across different environments.

## 📌 Core Concepts

### 1. What is a Module?

A module is a collection of related Go packages that are versioned together as a single unit. A module is defined by a `go.mod` file at its root.

### 2. Key Commands

- `go mod init <module-name>`: Initializes a new module.
- `go get <package-url>`: Adds a new dependency to your module.
- `go mod tidy`: Removes unused dependencies and adds missing ones.
- `go mod verify`: Checks that dependencies have not been tampered with.
- `go mod vendor`: Creates a `vendor` directory containing all dependencies.

### 3. File Structure

- **`go.mod`**: Defines the module's path, Go version, and dependencies.
- **`go.sum`**: Contains expected cryptographic hashes of the dependencies to ensure integrity.

---

## 📂 Module Structure

- **[go-mod-basics](./go-mod-basics)**: Creating a module and using external dependencies.
- **[multi-module-project](./multi-module-project)**: Working with multiple local modules using `replace` or Go Workspaces.

---

## 🚀 Quick Start: Creating a Module

### 1. Initialize the Module

```bash
go mod init example.com/my-project
```

### 2. Add a Dependency

```bash
go get github.com/fatih/color
```

### 3. Use the Dependency

```go
package main

import "github.com/fatih/color"

func main() {
    color.Cyan("Hello, Go Modules!")
}
```

---

## 💡 Best Practices

1. **Semantic Versioning**: Go uses semantic versioning (SemVer) for modules.
2. **Commit `go.mod` and `go.sum`**: Always include these files in your version control.
3. **Use `go mod tidy`**: Keep your dependency files clean and up to date.
4. **Avoid Global Packages**: Unlike some other languages, Go handles dependencies per module, not globally.

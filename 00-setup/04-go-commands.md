# Go Commands

A quick cheat sheet of commonly used Go commands for everyday development.

## Run a Go Program
```bash
go run main.go
```
- Compiles and runs the program in one step.
- Useful for quick testing.

## Build a Go Program
```bash
go build
```
- Compiles the project and creates an executable.
- No output means build was successful.

### Build with custom name:
```bash
go build -o app.exe
```

## Initialize a Module
```bash
go mod init module-name
```
- Creates a go.mod file.
- Required for dependency management.

## Manage Dependencies
```bash
go mod tidy
```
- Adds missing dependencies.
- Removes unused dependencies.

```bash
go get github.com/package/name
```
- Downloads and adds a dependency.

## Format Go Code
```bash
go fmt ./...
```
- Formats all Go files in the project using gofmt.

## Run Tests
```bash
go test
```
- Runs all tests in the current package.
- Run all tests:

```bash
go test ./...
```
- Run with verbose output:

```bash
go test -v
```

## Check Environment
```bash
go env
```
- Displays Go environment variables.

## Clean Build Cache
```bash
go clean
```
- Removes build cache and temporary files.

## List Packages
```bash
go list ./...
```
- Lists all packages in the module.

## Summary

▶ go run → run code
🏗 go build → build binary
📦 go mod → manage modules
🎨 go fmt → format code
🧪 go test → test code
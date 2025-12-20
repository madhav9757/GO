# Go Workspace

This section explains how Go organizes projects using **modules** and how the workspace works.

## GOPATH (Old Way)
- Earlier, Go required all projects to live inside a single directory called `GOPATH`.
- Typical structure:
   ```bash
    GOPATH/
    ├── src/
    ├── pkg/
    └── bin/
    ```
- This approach is **mostly deprecated** and not recommended for new projects.

## Go Modules (Modern Way ✅)
- Go now uses **modules**, introduced in Go 1.11.
- A module is any folder that contains a `go.mod` file.
- You can create a Go project **anywhere on your system**.

### Create a New Module
```bash
go mod init github.com/yourname/project-name
```
- This creates a go.mod file that:

- Tracks dependencies

- Defines the module name

- Manages versions automatically

### Recommended Workspace Setup

Example project structure:
```bash
GO/
├── 00-setup/
├── 01-basics/
├── 02-control-flow/
└── go.mod
```
- Useful Commands
```bash
go mod tidy    # Cleans up unused dependencies
go list -m     # Lists modules
go env         # Shows Go environment variables
 ```
### Summary

- ❌ GOPATH is old and mostly unused
- ✅ Go Modules are the standard
- 📦 Each project should have its own go.mod
- 📁 Projects can live anywhere on your system
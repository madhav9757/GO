# Environment Variables

This section explains how to set Go-related environment variables on **Windows** and how to use `go env`.

## GOPATH
- `GOPATH` is the workspace directory used by Go.
- With Go modules, you usually **don’t need to set it manually**.
- Default GOPATH on Windows:
    ```bash
    C:\Users<username>\go
    ```
   
## PATH (Very Important)
Go must be added to the system `PATH` so you can run `go` commands from anywhere.

### Default Go Install Path
```bash
C:\Program Files\Go\bin
```

### Check if Go is in PATH
Open **Command Prompt** and run:
```bash
go version
```
If it works, `PATH`is already set correctly.

### Add Go to PATH Manually (If Needed)

1. Press `Win + R` → type `sysdm.cpl` → Enter
2. Go to Advanced → Environment Variables
3. Under System variables, select `Path` → Edit
4. Click New and add:
    ```bash
    C:\Program Files\Go\bin
    ```
5. Click OK and restart your terminal
   
## Using go env
The `go env` command shows Go’s environment configuration.
```bash
go env
```
Check specific variables:
```bash
go env GOPATH
go env GOROOT
```

## Common Go Environment Variables

- `GOROOT` → Go installation directory
- `GOPATH` → Workspace directory
- `GOOS` → Target operating system
- `GOARCH` → Target architecture
Example:
    ```bash
    go env GOOS GOARCH
    ```

## Summary

- ✅ PATH must include Go binary
- ⚠ GOPATH usually not needed manually
- 🔍 go env helps debug environment issues
- 🚀 Ready to work with Go on Windows
# Basics of File Handling

This section covers the fundamental operations of reading from and writing to files in Go.

## 📖 What's Covered

1. **Reading Files**:
   - Using `os.ReadFile` for small files.
   - Using `os.Open` and `bufio.Scanner` for line-by-line reading.
   - Using `os.Open` and `io.Reader` for buffered reading.

2. **Writing Files**:
   - Using `os.WriteFile` for simple writes.
   - Using `os.Create` and `file.WriteString`.
   - Appending to existing files using `os.OpenFile`.

## 🛠️ Examples

- `reading/reading.go`: Demonstrates various ways to read content.
- `writing/writing.go`: Demonstrates creating, writing, and appending to files.

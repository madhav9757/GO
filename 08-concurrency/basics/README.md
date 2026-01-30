# Goroutines

Goroutines are functions or methods that run concurrently with other functions or methods. They are lightweight threads managed by the Go runtime.

## 📌 Key Concepts

- **Lightweight**: You can easily run thousands of goroutines. They start with a tiny stack (kB) that grows as needed.
- **Concurrency vs Parallelism**: Go provides concurrency (structure of the program). Parallelism (execution) depends on the available CPU cores.
- **`go` Keyword**: Prefix a function call with `go` to run it in a new goroutine.

## ⚠️ The Main Function

If the `main` function returns, the program exits immediately, and all other running goroutines are terminated. You must ensure the main function waits for others to finish (usually using `sync.WaitGroup` or Channels).

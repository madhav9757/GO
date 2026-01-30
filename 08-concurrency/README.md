# 08 - Concurrency

Concurrency is one of Go's standout features. Unlike other languages that rely on OS threads, Go uses **Goroutines**—lightweight threads managed by the Go runtime.

## 📂 Module Structure

This module covers the following topics:

### 1. [Basics](./basics)
**`basics/goroutines.go`**
- Starting goroutines with `go`.
- Synchronous vs Asynchronous execution.
- Anonymous goroutines.

### 2. [Channels](./channels)
**`channels/channels.go`**, **`channels/buffered-channels.go`**
- Communicating between goroutines.
- Buffered vs Unbuffered channels.
- `select` statement for multiplexing.
- Closing channels.

### 3. [Synchronization](./sync)
**`sync/waitgroups.go`**, **`sync/mutex.go`**
- `sync.WaitGroup`: Waiting for multiple goroutines.
- `sync.Mutex`: Safe access to shared memory (mutual exclusion).

### 4. [Patterns](./patterns)
**`patterns/worker-pool.go`**, **`patterns/context.go`**
- **Worker Pools**: distributing work across fixed workers.
- **Context**: Timed cancellation and deadline management.

## 🚀 Running the Examples

```bash
# Basics
go run basics/goroutines.go

# Channels
go run channels/channels.go
go run channels/buffered-channels.go

# Synchronization
go run sync/waitgroups.go
go run sync/mutex.go

# Patterns
go run patterns/worker-pool.go
go run patterns/context.go
```

## 🧠 Theory: Concurrency vs Parallelism

- **Concurrency**: Dealing with multiple things at once (Structure).
- **Parallelism**: Doing multiple things at once (Execution).

> "Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once." - Rob Pike

Go enables concurrency with **Goroutines** and **Channels**. The Go scheduler then maps these goroutines onto OS threads to achieve parallelism on multi-core processors.

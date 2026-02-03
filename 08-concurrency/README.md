# 08 - Concurrency

Concurrency is one of Go's standout features. Unlike other languages that rely on OS threads, Go uses **Goroutines**—lightweight threads managed by the Go runtime.

## 📂 Module Structure

This module covers the following topics:

### 1. [Basics](./basics)

**`basics/goroutines/goroutines.go`**

- Starting goroutines with `go`.
- Synchronous vs Asynchronous execution.
- Anonymous goroutines.

### 2. [Channels](./channels)

**`channels/basic/channels.go`**, **`channels/buffered/buffered-channels.go`**

- Communicating between goroutines.
- Buffered vs Unbuffered channels.
- `select` statement for multiplexing.
- Closing channels.

### 3. [Synchronization](./sync)

**`sync/waitgroup/waitgroups.go`**, **`sync/mutex/mutex.go`**

- `sync.WaitGroup`: Waiting for multiple goroutines.
- `sync.Mutex`: Safe access to shared memory (mutual exclusion).

### 4. [Patterns](./patterns)

**`patterns/worker-pool/worker-pool.go`**, **`patterns/context/context.go`**

- **Worker Pools**: distributing work across fixed workers.
- **Context**: Timed cancellation and deadline management.

## 🚀 Running the Examples

```bash
# Basics
go run basics/goroutines/goroutines.go

# Channels
go run channels/basic/channels.go
go run channels/buffered/buffered-channels.go

# Synchronization
go run sync/waitgroup/waitgroups.go
go run sync/mutex/mutex.go

# Patterns
go run patterns/worker-pool/worker-pool.go
go run patterns/context/context.go
```

## 🧠 Theory: Concurrency vs Parallelism

- **Concurrency**: Dealing with multiple things at once (Structure).
- **Parallelism**: Doing multiple things at once (Execution).

> "Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once." - Rob Pike

Go enables concurrency with **Goroutines** and **Channels**. The Go scheduler then maps these goroutines onto OS threads to achieve parallelism on multi-core processors.

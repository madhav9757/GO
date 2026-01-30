# Synchronization (sync)

While channels are the preferred way to communicate, sometimes you need lower-level primitives to synchronize memory access or execution flow.

## 📌 WaitGroup

`sync.WaitGroup` waits for a collection of goroutines to finish.

1.  `Add(n)`: Sets the number of goroutines to wait for.
2.  `Done()`: Decrements the counter (usually called via `defer` in the goroutine).
3.  `Wait()`: Blocks until the counter is 0.

## 🔒 Mutex (Mutual Exclusion)

`sync.Mutex` ensures that only one goroutine operates on a critical section of code (usually accessing shared memory) at a time.

- `Lock()`: Acquire the lock. Blocks if already locked.
- `Unlock()`: Release the lock.

**RWMutex**: If you have many readers but few writers, `sync.RWMutex` allows multiple concurrent readers but only one writer.

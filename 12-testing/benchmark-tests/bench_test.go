package benchmark

import "testing"

// BenchmarkFibonacci20 benchmarks the recursive Fibonacci function
func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

// BenchmarkFibonacciIterative20 benchmarks the iterative Fibonacci function
func BenchmarkFibonacciIterative20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(20)
	}
}

// BenchmarkFibonacciTable benchmarks with different inputs
func BenchmarkFibonacciTable(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"N=10", 10},
		{"N=20", 20},
		{"N=30", 30},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fibonacci(bm.n)
			}
		})
	}
}

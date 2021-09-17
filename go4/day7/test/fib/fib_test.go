package fib

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(2)
	}
}

//内部调用的函数
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

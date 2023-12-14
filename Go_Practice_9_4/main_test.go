package main

import (
	"testing"
)

func bench(b *testing.B, n int) {
	in, out := pipeline(n)
	for i := 0; i < b.N; i++ {
		go func() { in <- 42 }()
		<-out
	}
	close(in)
}

func BenchmarkPipeline1(b *testing.B) {
	bench(b, 1)
}

func BenchmarkPipeline2(b *testing.B) {
	bench(b, 2)
}

func BenchmarkPipeline4(b *testing.B) {
	bench(b, 4)
}

func BenchmarkPipeline8(b *testing.B) {
	bench(b, 8)
}

func BenchmarkPipeline16(b *testing.B) {
	bench(b, 16)
}

func BenchmarkPipeline32(b *testing.B) {
	bench(b, 32)
}

func BenchmarkPipeline64(b *testing.B) {
	bench(b, 64)
}

func BenchmarkPipeline128(b *testing.B) {
	bench(b, 128)
}

func BenchmarkPipeline256(b *testing.B) {
	bench(b, 256)
}

func BenchmarkPipeline512(b *testing.B) {
	bench(b, 512)
}

func BenchmarkPipeline1024(b *testing.B) {
	bench(b, 1024)
}

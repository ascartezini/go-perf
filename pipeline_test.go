package main

import (
	"testing"
)

func Benchmark_readablestream(b *testing.B) {

	for i := 0; i < b.N; i++ {
		readableStream()
	}
}

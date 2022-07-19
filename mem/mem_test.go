package main

import (
	"testing"
)

// func BenchmarkHeapAllocation(t *testing.B) {
// 	for i := 0; i < t.N; i++ {
// 		heapAlloc()
// 	}
// }

// func BenchmarkStackAllocation(t *testing.B) {
// 	for i := 0; i < t.N; i++ {
// 		stackAlloc()
// 	}
// }

// func BenchmarkHeapAllocation(t *testing.B) {
// 	for i := 0; i < t.N; i++ {
// 		heapUserAlloc()
// 	}
// }

// func BenchmarkStackAllocation(t *testing.B) {
// 	for i := 0; i < t.N; i++ {
// 		stackUserAlloc()
// 	}
// }

// func BenchmarkPasswordHash(t *testing.B) {
// 	r := Register{Password: "123456"}
// 	for i := 0; i < t.N; i++ {
// 		r.HashPasswd()
// 	}
// }

func BenchmarkFormatSprintf(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FormatSprintf()
	}
}

func BenchmarkFormatPlus(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FormatPlusSign()
	}
}

func BenchmarkFormatBuilder(t *testing.B) {
	for i := 0; i < t.N; i++ {
		FormatBuilder()
	}
}

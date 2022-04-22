package main

import (
	"testing"
)


func BenchmarkNonGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumNonGenerics(1234.1234, 9876.9876)
	}
}
func BenchmarkGenericsFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumGenerics(1234.1234, 9876.9876)
	}
}
package main

import (
	"testing"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strings()
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RegExp()
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Matcher()
	}
}

func BenchmarkD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuncMatcher()
	}
}

package main

import (
	"strings"
	"testing"
)

func concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func concatOptimized(str []string) string {
	var builder strings.Builder
	totalLength := 0
	for _, s := range str {
		totalLength += len(s)
	}
	builder.Grow(totalLength)
	for _, s := range str {
		builder.WriteString(s)
	}
	return builder.String()
}

func generateTestData(n int) []string {
	data := make([]string, n)
	for i := range data {
		data[i] = "test"
	}
	return data
}

func BenchmarkConcat(b *testing.B) {
	data := generateTestData(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = concat(data)
	}
}

func BenchmarkConcatOptimized(b *testing.B) {
	data := generateTestData(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = concatOptimized(data)
	}
}

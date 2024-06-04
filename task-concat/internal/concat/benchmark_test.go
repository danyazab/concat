package concat

import (
	"testing"
)

// OriginalConcat is the non-optimized version
func OriginalConcat(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}

func generateInput(size int) []string {
	input := make([]string, size)
	for i := 0; i < size; i++ {
		input[i] = "test"
	}
	return input
}

func BenchmarkOriginalConcat(b *testing.B) {
	input := generateInput(30)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OriginalConcat(input)
	}
}

func BenchmarkConcat(b *testing.B) {
	input := generateInput(30)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concat(input)
	}
}

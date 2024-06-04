package concat

import "strings"

// Concat is an optimized version
func Concat(strs []string) string {
	//Using strings.Builder allows you to avoid multiple allocation of memory
	var builder strings.Builder
	totalLength := 0
	for _, s := range strs {
		totalLength += len(s)
	}
	// Using Grow further reduces the number of memory allocation operations
	//because it allocates memory in advance
	builder.Grow(totalLength)
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

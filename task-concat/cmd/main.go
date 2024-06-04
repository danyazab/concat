package main

import (
	"fmt"
	"task-concat/internal/concat"
)

func main() {
	strings := []string{"Hi", "!", " ", "My", " ", "name", " ", "is", " ", "D", "a", "n", "i", "e", "l", "!"}
	result := concat.Concat(strings)
	fmt.Println(result)
}

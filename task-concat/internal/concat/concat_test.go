package concat

import "testing"

func TestConcat(t *testing.T) {
	const expected = "Hi! My name is Daniel!"

	input := []string{"Hi", "!", " ", "My", " ", "name", " ", "is", " ", "D", "a", "n", "i", "e", "l", "!"}
	result := Concat(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

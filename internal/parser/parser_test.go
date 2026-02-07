package parser

import "testing"

func TestExample(t *testing.T) {
	sum := 2 + 2
	if sum != 4 {
		t.Errorf("Expected 4, got %d", sum)
	}
}

func TestExample2(t *testing.T) {
	sum := 2 + 3
	if sum != 5 {
		t.Errorf("Expected 4, got %d", sum)
	}
}
func TestExample3(t *testing.T) {
	sum := 3 + 3
	if sum != 6 {
		t.Errorf("Expected 4, got %d", sum)
	}
}

package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Errorf("add(1, 2) = %d; want 3", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(1, 2)
	if result != -1 {
		t.Errorf("subtract(1, 2) = %d; want -1", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(1, 2)
	if result != 2 {
		t.Errorf("multiply(1, 2) = %d; want 2", result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(1, 2)
	if result != 0 {
		t.Errorf("divide(1, 2) = %d; want 0", result)
	}
}

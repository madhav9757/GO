package math

import "testing"

// Basic Test
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// Table-driven Test (The Go Way)
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 6},
		{"negative numbers", -2, -3, 6},
		{"zero", 10, 0, 0},
		{"one negative", -5, 4, -20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("%s: Multiply(%d, %d) = %d; want %d", tt.name, tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Subtests for Division
func TestDivide(t *testing.T) {
	t.Run("normal division", func(t *testing.T) {
		result := Divide(10, 2)
		if result != 5 {
			t.Errorf("Divide(10, 2) = %d; want 5", result)
		}
	})

	t.Run("divide by zero", func(t *testing.T) {
		result := Divide(10, 0)
		if result != 0 {
			t.Errorf("Divide(10, 0) = %d; want 0", result)
		}
	})
}

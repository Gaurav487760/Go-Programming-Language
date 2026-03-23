package mathutils

import "testing"

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{5, 3, 2},
		{10, 4, 6},
		{7, 7, 0},
		{0, 5, -5},
	}

	for _, tt := range tests {
		result := Subtract(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Subtract(%d, %d) = %d; expected %d",
				tt.a, tt.b, result, tt.expected)
		}
	}
}
package algorithms

import (
	"fmt"
	"testing"
)

func TestSigleNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{1, 2, 2, 3, 4, 1, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			out := singleNumber(tt.input)
			if out != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, out)
			}
		})
	}
}

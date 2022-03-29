package algorithms

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := majorityElement2(tt.input)
			if output != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, output)
			}
		})
	}
}

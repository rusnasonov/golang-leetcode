package algorithms

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			[]int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			[]int{},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			tree := buildTree(tt.input)
			inverted := invertTree(tree)
			values := treeValues(inverted)
			for i, v := range tt.expected {
				if v != values[i] {
					t.Errorf("expected %d, got %d", v, values[i])
				}
			}
		})
	}
}

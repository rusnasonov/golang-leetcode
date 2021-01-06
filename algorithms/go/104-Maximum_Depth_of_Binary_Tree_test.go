package main

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct{
		input []int
		expected int
	}{
		{[]int{3,9,20,-1,-1,15,7}, 3},
		{[]int{1,-1,2}, 2},
		{[]int{1,2,3,4,-1,-1,5}, 3},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			tree := buildTree(tt.input)
			depth := maxDepthRecursive(tree)
			if depth != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, depth)
			}
		})
	}
}
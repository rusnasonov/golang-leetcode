package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct{
		input []int
		expected []int
	}{
		{[]int{1,-1,2,3}, []int{1,3,2}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1,2}, []int{2,1}},
		{[]int{1,-1,2}, []int{1,2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			tree := buildTree(tt.input)
			values := inorderTraversal(tree)
			if !reflect.DeepEqual(values, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, values)
			}
		})
	}

}
package algorithms

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct{
		input struct{
			tree []int
			k int
		}
		expexted int
	}{
		{
			struct{tree []int; k int}{[]int{3,1,4,-1,2}, 1}, 1,
		},
		{
			struct{tree []int; k int}{[]int{5,3,6,2,4,-1,-1,1}, 3}, 3,
		},
		{
			struct{tree []int; k int}{[]int{1}, 1}, 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input.tree), func(t *testing.T) {
			output := kthSmallestChannel(
				buildTree(tt.input.tree),
				tt.input.k,
			)
			if output != tt.expexted {
				t.Errorf("expected %d, got %d", tt.expexted, output)
			}
		})
	}
}

func BenchmarkKthSmallest(b *testing.B) {
	tree := buildTree([]int{5,3,6,2,4,-1,-1,1})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kthSmallest(tree, 4)
	}
}

func BenchmarkKthSmallestRecursive(b *testing.B) {
	tree := buildTree([]int{5,3,6,2,4,-1,-1,1})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kthSmallestRecursive(tree, 4)
	}
}

func BenchmarkKthSmallestChannel(b *testing.B) {
	tree := buildTree([]int{5,3,6,2,4,-1,-1,1})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kthSmallestChannel(tree, 4)
	}
}
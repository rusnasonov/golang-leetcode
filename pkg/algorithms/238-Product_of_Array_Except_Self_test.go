package algorithms

import (
	"fmt"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := productExceptSelf(tt.input)
			for i, v := range tt.expected {
				if v != output[i] {
					t.Errorf("expected %v, got %v", v, output[i])
				}
			}
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	t := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		t[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productExceptSelf(t)
	}
}

func BenchmarkProductExceptSelf2(b *testing.B) {
	t := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		t[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		productExceptSelf2(t)
	}
}

package algorithms

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		input struct {
			arr []int
			k   int
		}
		output []int
	}{
		{
			struct {
				arr []int
				k   int
			}{[]int{1, 1, 1, 2, 2, 3}, 2},
			[]int{1, 2},
		},
		{
			struct {
				arr []int
				k   int
			}{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6}, 6},
			[]int{6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input.arr), func(t *testing.T) {
			output := topKFrequent(tt.input.arr, tt.input.k)
			if !reflect.DeepEqual(tt.output, output) {
				t.Errorf("expected %v, got %v", tt.output, output)
			}
		})
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 10000; i++ {
		data[i] = rand.Intn(2000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topKFrequent(data, 500)
	}
}

func BenchmarkTopKFrequentHeap(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 10000; i++ {
		data[i] = rand.Intn(2000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		topKFrequentHeap(data, 500)
	}
}

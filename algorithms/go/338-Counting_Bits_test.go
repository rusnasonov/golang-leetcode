package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			r := countBitsDynamic(tt.input)
			if !reflect.DeepEqual(r, tt.expected) {
				t.Errorf("expected: %v, got %v", tt.expected, r)
			}
		})
	}
}

func BenchmarkCountBitsMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countBits(10000)
	}
}

func BenchmarkCountBitsDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countBitsDynamic(10000)
	}
}

func BenchmarkCountBitsShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countBitsShift(10000)
	}
}
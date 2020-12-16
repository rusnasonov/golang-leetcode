package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct{
		input []int
		expected [][]int
	}{
		{[]int{1,2,3}, [][]int{{},{1},{2},{1,2},{3},{1,3},{2,3},{1,2,3}}},
		{[]int{0}, [][]int{{},{0}}},
		{[]int{}, [][]int{{}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := subsets(tt.input)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
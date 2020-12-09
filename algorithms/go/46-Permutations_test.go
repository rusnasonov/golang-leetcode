package main

import (
	"fmt"
	"testing"
	"reflect"
	"sort"
)

func TestPermute(t *testing.T) {
	tests := []struct{
		input []int
		expected [][]int
	}{
		{[]int{1,2,3}, [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			r := permute(tt.input)
			sort.Slice(r, func(i, j int) bool {
				if r[i][0] == r[j][0] {
					return r[i][1] < r[j][1]
				} else {
					return r[i][0] < r[j][0]
				}
			})
			if !reflect.DeepEqual(r, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, r)
			}
		})
	}
}

func BenchmarkIterativePermute(b *testing.B) {
	for i:=0; i < b.N; i++ {
		permute([]int{1,2,3,4,5,6,7,8,9,10})
	}
}

func BenchmarkRecursivePermute(b *testing.B) {
	for i:=0; i < b.N; i++ {
		permuteRecursive([]int{1,2,3,4,5,6,7,8,9,10})
	}
}
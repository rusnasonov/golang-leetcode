package main

import (
	"reflect"
	"testing"
)

func TestTreeValues(t *testing.T) {
	cases := [][]int{
		{2,1,3,-1,4,-1,7},
		{1,3,2,5,-1},
	}
	for _, c := range cases {
		tree := buildTree(c)
		values := treeValues(tree)
		if !reflect.DeepEqual(c, values) {
			t.Errorf("expected %v, got %v", c, values)
		}
	}
}

func TestMergeTrees(t *testing.T) {
	merged := mergeTrees(
		buildTree([]int{1,3,2,5,-1}),
		buildTree([]int{2,1,3,-1,4,-1,7}),
	)
	values := treeValues(merged)
	if !reflect.DeepEqual(values, []int{3,4,5,5,4,-1,7}) {
		t.Errorf("expected %v, got %v", []int{3,4,5,5,4,-1,7}, values)
	}
}
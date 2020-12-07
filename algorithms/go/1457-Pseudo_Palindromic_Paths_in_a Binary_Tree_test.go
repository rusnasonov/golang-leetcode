package main

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	root := buildTree([]int{2, 3, 1, 3, 1, -1, 1})
	if root.Val != 2 || root.Left.Val != 3 || root.Right.Val != 1 {
		t.Error("bad root node")
	}
}

func TestIsPalindromic(t *testing.T) {
	var tests = []struct {
		input    []int
		expected bool
	}{
		{[]int{1}, true},
		{[]int{1, 1}, true},
		{[]int{1, 1, 2}, true},
		{[]int{1, 1, 2, 2}, true},
		{[]int{1, 1, 2, 2, 3}, true},
		{[]int{1, 1, 2, 2, 3, 4}, false},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3}, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := isPalindromic(tt.input)
			if ans != tt.expected {
				t.Errorf("expected %v, got %v", ans, tt.expected)
			}
		})
	}
}

func TestPseudoPalindromicPaths(t *testing.T) {
	var tests = []struct {
		input    []int
		expexted int
	}{
		{[]int{2, 1, 1, 1, 3, -1, -1, -1, -1, -1, 1}, 1},
		{[]int{2, 3, 1, 3, 1, -1, 1}, 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.input)
		t.Run(testname, func(t *testing.T) {
			tree := buildTree(tt.input)
			count := pseudoPalindromicPaths(tree)
			if count != tt.expexted {
				t.Errorf("expected %d, got %d", tt.expexted, count)
			}
		})
	}

}

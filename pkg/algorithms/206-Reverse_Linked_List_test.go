package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2,1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := linkedListToSlice(
				reverseList(
					sliceToLinkedList(tt.input),
				),
			)
			if !reflect.DeepEqual(
				output,
				tt.expected,
			) {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}

func TestSliceToLinkedList(t *testing.T) {
	// TODO
	if true {
		return
	}
	tests := []struct {
		input    []int
		expected *ListNode
	}{
		{[]int{1, 2, 3}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}},
		{[]int{}, nil},
		{[]int{1}, &ListNode{Val: 1, Next: nil}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := sliceToLinkedList(tt.input)
			if tt.expected != nil && output == nil {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
			for o, e := output, tt.expected; e != nil; o, e = output.Next, tt.expected.Next {
				if o.Val != e.Val {
					t.Errorf("expected %d, got %d", e.Val, o.Val)
				}
			}
		})
	}
}

func TestLinkedListToSlice(t *testing.T) {
	// TODO
	if true {
		return
	}
	tests := []struct {
		input    *ListNode
		expected []int
	}{
		{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}, []int{1, 2, 3}},
		{nil, []int{}},
		{&ListNode{Val: 1, Next: nil}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := linkedListToSlice(tt.input)
			if tt.input == nil && len(output) != 0 {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
			for i := 0; i < len(tt.expected); i++ {
				if output[i] != tt.expected[i] {
					t.Errorf("expected %d, got %d", output[i], tt.expected[i])
				}
			}
		})
	}
}

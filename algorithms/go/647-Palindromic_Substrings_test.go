package main

import (
	"testing"
)

func TestCountSubstring(t *testing.T) {
	tests := []struct{
		input string
		expected int
	}{
		{"abc", 3},
		{"aaa", 6},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := countSubstringsDynamic(tt.input)
			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct{
		input string
		expected bool
	}{
		{"abc", false},
		{"aaa", true},
		{"abccba", true},
		{"a", true},
		{"abcba", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := isPalindrome(&tt.input, 0, len(tt.input))
			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
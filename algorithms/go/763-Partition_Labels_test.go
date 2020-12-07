package main

import (
	"reflect"
	"testing"
)

func TestPartionLabels(t *testing.T) {
	tests := []struct{
		input string
		expected []int
	}{
		{"ababcbacadefegdehijhklij", []int{9,7,8}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := partitionLabels(tt.input)
			if !reflect.DeepEqual(r, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, r)
			}
		})
	}
}

func TestGetIntervals(t *testing.T) {
	intervals := getIntervals("ababcbacadefegdehijhklij")
	expected := map[rune]interval{
		'a': {'a',0,8},
		'b': {'b',1,5},
		'c': {'c',4,7},
		'd': {'d',9,14},
		'e': {'e',10,15},
		'f': {'f',11,11},
		'g': {'g',13,13},
		'h': {'h',16,19},
		'i': {'i',17,22},
		'j': {'j',18,23},
		'l': {'k',21,21},
		'k': {'l',20,20},
	}
	for k, v := range intervals {
		if expected[k].start != v.start || expected[k].end != v.end {
			t.Errorf("'%s'. expected %v, got '%v'", string(k), expected[k], v)
		}
	}
}
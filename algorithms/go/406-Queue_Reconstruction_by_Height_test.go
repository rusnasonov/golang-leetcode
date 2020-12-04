package main

import (
	"testing"
	"fmt"
	"reflect"
)


func TestCheck(t *testing.T) {
	if !check([][]int{{5,0},{7,0},{5,2},{6,1},{4,4},{7,1}}) {
		t.Errorf("expected true, got false")
	}
	if check([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}) {
		t.Errorf("expected false, got true")
	}
	if check([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}) {
		t.Errorf("expected false, got true")
	}
}

func TestReconstructQueue(t *testing.T) {
	var tests = []struct{
		input [][]int
		expected [][]int
	}{
		{
			[][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}},
			[][]int{{5,0},{7,0},{5,2},{6,1},{4,4},{7,1}},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.input)
        t.Run(testname, func(t *testing.T) {
            ans := reconstructQueue(tt.input)
            if !reflect.DeepEqual(ans, tt.expected) {
                t.Errorf("expected %v, got %v", tt.expected, ans)
            }
        })
	}
}

func BenchmarkReconstructQueue(b *testing.B) {
    for i := 0; i < b.N; i++ {
        reconstructQueue([][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}})
    }
}

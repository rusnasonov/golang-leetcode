package main

import (
	"testing"
	"fmt"
)

type intervals [][]int

type input struct{
	a intervals
	b intervals
}

func TestIntervalListIntersection(t *testing.T) {
	tests := []struct{
		input 
		expexted intervals 
	}{
		{
			input{
				intervals{{0,2},{5,10},{13,23},{24,25}},
				intervals{{1,5},{8,12},{15,24},{25,26}},
			},
			intervals{{1,2},{5,5},{8,10},{15,23},{24,24},{25,25}},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.input)
		t.Run(testname, func(t *testing.T){
			output := intervalIntersection(tt.a, tt.b)
			for i, item := range output {
				if item[0] != tt.expexted[i][0] || item[1] != tt.expexted[i][1]{
					t.Errorf("expected %v, got %v", tt.expexted[i], item)
				}
			}
			
		})
	}
}

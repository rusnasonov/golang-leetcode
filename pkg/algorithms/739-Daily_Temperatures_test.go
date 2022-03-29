package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDailyTempratures(t *testing.T) {
	tests := []struct {
		input    []int
		expexted []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			output := dailyTemperatures(tt.input)
			if !reflect.DeepEqual(output, tt.expexted) {
				t.Errorf("expected %v, got %v", tt.expexted, output)
			}
		})
	}
}

func generateTest(N int) []int {
	res := make([]int, N)
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			res[i] = 71
		} else {
			res[i] = 76
		}
	}
	return res
}

func BenchmarkDailyTemperatures(b *testing.B) {
	data := generateTest(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dailyTemperatures(data)
	}
}

func BenchmarkDailyTemperaturesStack(b *testing.B) {
	data := generateTest(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dailyTemperaturesStack(data)
	}
}

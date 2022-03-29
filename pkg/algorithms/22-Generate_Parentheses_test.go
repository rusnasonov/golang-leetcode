package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		input    int
		expected []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			output := generateParenthesis(tt.input)
			if !reflect.DeepEqual(tt.expected, output) {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}

}

func BenchmarkGenerateParenthesisIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateParenthesis(10)
	}
}

func BenchmarkGenerateParenthesisBacktrack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateParenthesisBacktrack(10)
	}
}

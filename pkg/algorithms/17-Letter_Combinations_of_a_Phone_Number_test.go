package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{"2", []string{"a", "b", "c"}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"237", []string{
			"adp", "adq", "adr", "ads", "aep", "aeq", "aer", "aes", "afp", "afq", "afr", "afs",
			"bdp", "bdq", "bdr", "bds", "bep", "beq", "ber", "bes", "bfp", "bfq", "bfr", "bfs",
			"cdp", "cdq", "cdr", "cds", "cep", "ceq", "cer", "ces", "cfp", "cfq", "cfr", "cfs",
		}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := letterCombinations(tt.input)
			if !reflect.DeepEqual(ans, tt.expected) {
				t.Errorf("got %v, expected %v", ans, tt.expected)
			}
		})
	}
}

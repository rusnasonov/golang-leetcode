package algorithms

import (
	"fmt"
	"math/bits"
	"strings"
)

func checkParenthesis(expr string) bool {
	open, close := 0, 0
	for _, q := range expr {
		if q == '(' {
			open++
		} else {
			close++
		}
		if open - close < 0 {
			return false
		}
	}
	return true
}

func generateParenthesis(n int) []string {
	q := 1 << (n*2-1)
	res := []string{}
	//threashold := ((1 << (n*2)) - 1) / 3 // It's binary sequence 010101... https://oeis.org/A002450
	for i := 0; i <= q; i++ {
		if (i | 1 == i) && (i | q != i) && bits.OnesCount(uint(i)) == n {
			s := fmt.Sprintf(fmt.Sprintf("%%0%db", n*2), i)
			s = strings.ReplaceAll(
				strings.ReplaceAll(s, "0", "("),
				"1",
				")",
			)
			if checkParenthesis(s) {
				res = append(res, s)
			}
		}
	}
	return res
}

func generateParenthesisBacktrack(n int) []string {
	res := []string{}
	backtrack(&res, "", 0, 0, n)
	return res
}

func backtrack(ans *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	if (open < max) {
		backtrack(ans, cur+"(", open+1, close, max)
	}
	if (close < open) {
		backtrack(ans, cur+")", open, close+1, max)
	}
}
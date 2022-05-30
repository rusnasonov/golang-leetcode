package algorithms

import "testing"

func TestDivideTwoInteger(t *testing.T) {
    cases := []struct{
        divident, divisor, expected int
    }{
        {10, 3, 3},
        {7, -3, -2},
        {-2147483648, -1, 2147483647},
        {-2147483648, 1, -2147483648},
        {2147483647, 3, 715827882},
    }
    for _, v := range cases {
        out := divide(v.divident, v.divisor)
        if out != v.expected {
            t.Errorf("got %d, expected %d", out, v.expected)
        }
    }
}

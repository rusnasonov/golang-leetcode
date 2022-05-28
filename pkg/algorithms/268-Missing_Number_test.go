package algorithms

import "testing"


func TestMissingNumber(t *testing.T) {
    cases := []struct{
        inp []int
        expected int
    }{
        {[]int{3,0,1}, 2},
        {[]int{0,1}, 2},
        {[]int{9,6,4,2,3,5,7,0,1}, 8},
    }
    for _, v := range cases {
        s := missingNumber(v.inp) 
        if s != v.expected {
            t.Errorf("expected %d, got %d", v.expected, s)
        }
    }
}

package algorithms

import "testing"

func sliceEq[T comparable](left, right []T) bool {
    if len(left) != len(right) {
        return false
    }
    for i := 0; i < len(left); i++ {
        if left[i] != right[i] {
            return false
        }
    }
    return true
}

func TestTwoSum(t *testing.T) {
    cases := []struct{
        nums []int
        target int
        expected []int
    }{
        {[]int{2,7,11,15}, 9, []int{1,0}},
        {[]int{3,2,4}, 6, []int{2,1}},
    }
    for _, v := range cases {
        got := twoSum(v.nums, v.target)
        if !sliceEq(got, v.expected) {
            t.Errorf("got %v, expected %v", got, v.expected)
        }
    }
}

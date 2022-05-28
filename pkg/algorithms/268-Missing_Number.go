package algorithms

func missingNumber(nums []int) int {
    n := len(nums)
    s1 := (n*(n+1))/2
    s2 := 0
    for _, v := range nums {
        s2 = s2 + v
    }
    return s1 - s2 
}

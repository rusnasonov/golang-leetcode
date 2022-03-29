package algorithms

func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		lenR := len(res)
		curElem := nums[i]
		for j := 0; j < lenR; j++ {
			s := []int{}
			currSet := res[j]
			if len(currSet) > 0 && curElem > currSet[len(currSet)-1] {
				s = append(append(res[j][:0:0], currSet...), curElem)
			} else {
				s = append([]int{curElem}, currSet...)
			}
			res = append(res, s)
		}
	}
	return res
}

// Simple, but other order
func subsets2(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			res = append(res, nums[i:j+1])
		}
	}
	return res
}
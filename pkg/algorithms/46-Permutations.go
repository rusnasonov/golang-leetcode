package algorithms

// See https://en.wikipedia.org/wiki/Heap%27s_algorithmpa
func permute(nums []int) [][]int {
	res := [][]int{}
	c := make([]int, len(nums))
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	res = append(res, tmp)
	i := 0
	for i < len(nums) {
		if c[i] < i {
			if i%2 == 0 {
				nums[0], nums[i] = nums[i], nums[0]
			} else {
				nums[c[i]], nums[i] = nums[i], nums[c[i]]
			}
			tmp = make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return res
}

func permuteRecursive(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		v := make([]int, len(nums))
		copy(v, nums)
		v[i], v[0] = v[0], v[i]
		v = v[1:]
		r := permuteRecursive(v)
		for _, j := range r {
			j = append([]int{nums[i]}, j...)
			res = append(res, j)
		}
	}
	return res
}

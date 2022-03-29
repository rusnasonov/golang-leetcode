package algorithms

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	left := make([]int, len(nums)+1)
	right := make([]int, len(nums)+1)
	left[0] = 1
	right[0] = 1
	for i := 0; i < len(nums); i++ {
		left[i+1] = left[i] * nums[i]
		right[i+1] = right[i] * nums[len(nums)-1-i]
	}
	for i := 0; i < len(nums); i++ {
		out[i] = left[i] * right[len(nums)-i-1]
	}
	return out
}

// Reduce Memory Сonsumption
func productExceptSelf2(nums []int) []int {
	lenght := len(nums)
	out := make([]int, lenght)
	out[0] = 1
	for i := 1; i < lenght; i++ {
		out[i] = nums[i-1] * out[i-1]
	}
	r := 1
	for i := lenght - 1; i >= 0; i-- {
		out[i] = out[i] * r
		r *= nums[i]
	}
	return out
}

package main

func singleNumber(nums []int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}
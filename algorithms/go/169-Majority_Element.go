package main


// O(n) time, O(n/2) space
func majorityElement(nums []int) int {
	counter := map[int]int{}
	max := 0
	elem := 0
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
		if counter[nums[i]] > max {
			max = counter[nums[i]]
			elem = nums[i]
		}
	}
	return elem
}

func majorityElement2(nums []int) int {
	count := 0
	candidate := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
package algorithms

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	watched := map[int][]int{}
	for i := 0; i < len(T); i++ {
		currTemp := T[i]
		watched[currTemp] = append(watched[currTemp], i)
		for watchedTemp, indexes := range watched {
			if currTemp > watchedTemp {
				for _, idx := range indexes {
					res[idx] = i - idx
				}
				delete(watched, watchedTemp)
			}
		}
	}
	return res
}

func dailyTemperaturesStack(T []int) []int {
	result := make([]int, len(T))
	stack := make([]int, 0, len(T)/2)

	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}

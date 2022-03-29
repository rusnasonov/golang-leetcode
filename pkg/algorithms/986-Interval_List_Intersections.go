package algorithms

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	for i, j := 0, 0; i < len(A) && j < len(B); {
		a := A[i]
		b := B[j]
		if a[1] < b[0] {
			i++
			continue
		}
		if b[1] < a[0] {
			j++
			continue
		}
		start := max(a[0], b[0])
		end := min(a[1], b[1])
		res = append(res, []int{start, end})
		if a[1] < b[1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

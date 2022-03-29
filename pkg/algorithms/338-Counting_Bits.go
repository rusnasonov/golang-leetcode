package algorithms

import "math/bits"

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = bits.OnesCount64(uint64(i))
	}
	return res
}

func countBitsDynamic(num int) []int {
	bitTable := make([]int, num+1)
	bitTable[0] = 0
	q := 2
	for i := 1; i <= num; i++ {
		bitTable[i] = bitTable[i%q] + 1
		if i%q == 0 {
			q = i
		}
	}
	return bitTable
}

func countBitsShift(num int) []int {
	counts := make([]int, num+1)

	for i := 1; i <= num; i++ {
		count := counts[(i>>1)] + (i & 1)
		counts[i] = count
	}

	return counts
}

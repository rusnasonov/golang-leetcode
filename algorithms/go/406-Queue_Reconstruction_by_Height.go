// 406. Queue Reconstruction by Height
// https://leetcode.com/problems/queue-reconstruction-by-height/
// Medium

package main

import (
	"sort"
)

type queue [][]int

func (self queue) Len() int {return len(self)}
func (self queue) Swap(i, j int) { self[i], self[j] = self[j], self[i] }
func (self queue) Less(i, j int) bool {
	if self[i][0] == self[j][0] {
		return self[i][1] > self[j][1]
	}
	return self[i][0] < self[j][0]
}

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	sort.Sort(queue(people))
	for _, item := range people {
		i := 0
		for emptyCells := -1; emptyCells != item[1]; {
			if len(res[i]) == 0 {
				emptyCells++
			}
			i++
		}
		res[i-1] = item
	}
    return res
}

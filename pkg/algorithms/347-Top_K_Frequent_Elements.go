package algorithms

import (
	"container/heap"
	"sort"
)

type kv struct {
	k int
	v int
}

type IntHeap []kv

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].v > h[j].v }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// First soltion
func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	output := make([]int, k)
	sortedCounter := []kv{}
	for _, num := range nums {
		counter[num]++
	}
	for k, v := range counter {
		sortedCounter = append(sortedCounter, kv{k, v})
	}
	sort.Slice(sortedCounter, func(i int, j int) bool {
		return sortedCounter[i].v > sortedCounter[j].v
	})
	for i := 0; i < k; i++ {
		output[i] = sortedCounter[i].k
	}
	return output
}

// Leetcode solution
func topKFrequentHeap(nums []int, k int) []int {
	counter := map[int]int{}
	output := make([]int, k)
	for _, num := range nums {
		counter[num]++
	}
	i := 0
	counterHeap := make(IntHeap, len(counter))
	for k, v := range counter {
		counterHeap[i] = kv{k, v}
		i++
	}
	heap.Init(&counterHeap)
	for i := 0; i < k; i++ {
		item := heap.Pop(&counterHeap).(kv)
		output[i] = item.k
	}
	return output
}

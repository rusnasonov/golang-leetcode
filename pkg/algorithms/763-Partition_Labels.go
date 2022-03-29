package algorithms

import "sort"

type interval struct {
	letter rune
	start int
	end int
}

func getIntervals(S string) map[rune]interval{
	res := make(map[rune]interval)
	for i, letter := range S {
		letterInterval, ok := res[letter]
		if !ok {
			res[letter] = interval{letter, i, i}
		} else {
			res[letter] = interval{letter, letterInterval.start, i}
		}
	}
	return res
}

func partitionLabels(S string) []int {
	var intervals []interval
	var res []int
	for _, v := range getIntervals(S) {
		intervals = append(intervals, v)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	start, end := intervals[0].start, intervals[0].end
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if curr.start > end {
			res = append(res, end-start+1)
			start = curr.start
			end = curr.end
			continue
		}
		if curr.end > end {
			end = curr.end
		}
	}
	res = append(res, end-start+1)
    return res
}
from typing import List

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        out = []
        intervals = sorted(intervals, key=lambda x: x[0])
        start, end = intervals[0]
        for s, e in intervals:
            if s <= end:
                end = max(end, e)
                start = min(start, s)
            else:
                out.append([start, end])
                start, end = s, e
        out.append([start, end])
        return out

def test_merge_intervals():
    assert Solution().merge([[1,3],[2,6],[8,10],[15,18]]) == [[1,6],[8,10],[15,18]]
    assert Solution().merge([[1,4],[4,5]]) == [[1,5]]
    assert Solution().merge([[1,4],[0,4]]) == [[0,4]]
    assert Solution().merge([[1,4],[0,0]]) == [[0,0],[1,4]]
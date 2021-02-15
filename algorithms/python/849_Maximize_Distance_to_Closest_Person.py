from typing import List

class Solution:
    def maxDistToClosest(self, seats: List[int]) -> int:
        max_dist = 0
        start, end = 0, len(seats) - 1
        while seats[start] == 0:
            start += 1
        max_dist = max(max_dist, start)
        while seats[end] == 0:
            end -= 1
        max_dist = max(max_dist, len(seats) - 1 - end)
        empty_seats_seq_len = 0
        while start <= end:
            if seats[start] == 0:
                empty_seats_seq_len += 1
            else:
                max_dist = max(max_dist, (empty_seats_seq_len+1)//2)
                empty_seats_seq_len = 0
            start += 1
        return max_dist


def test_maxDistToClosest():
    assert Solution().maxDistToClosest([1,0,0,0,1,0,1]) == 2
    assert Solution().maxDistToClosest([1,0,0,0]) == 3
    assert Solution().maxDistToClosest([0,0,0,1]) == 3
    assert Solution().maxDistToClosest([0,1,0,1,0]) == 1
    assert Solution().maxDistToClosest([0,1,0,0,1,0]) == 1
from typing import List

class Solution:
    def canAttendMeetings(self, intervals: List[List[int]]) -> bool:
        intervals = sorted(intervals, key=lambda x: x[0])
        for i in range(len(intervals)-1):
            if intervals[i][1] > intervals[i+1][0]:
                return False
        return True

def test_canAttendMeetings():
    assert Solution().canAttendMeetings([[1,2],[3,4],[7,8],[8,9]]) is True
    assert Solution().canAttendMeetings([[1,3],[2,4],[5,6]]) is False
    assert Solution().canAttendMeetings([[1,3]]) is True
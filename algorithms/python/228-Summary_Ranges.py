from typing import List


class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        #nums = sorted(nums)
        res = []
        start, end = None, None
        i = 0
        while i < len(nums):
            start = nums[i]
            while i < len(nums) - 1 and nums[i] + 1 == nums[i+1]:
                i+=1
            end = nums[i]
            if end == start:
                res.append(str(start))
            else:
                res.append(f"{start}->{end}")
            end = None
            i+=1
        return res



def test_summary_ranges():
    assert Solution().summaryRanges([0,1,2,4,5,7]) == ["0->2","4->5","7"]
    assert Solution().summaryRanges([0,2,3,4,6,8,9]) == ["0","2->4","6","8->9"]
    assert Solution().summaryRanges([]) == []
    assert Solution().summaryRanges([-1]) == ["-1"]


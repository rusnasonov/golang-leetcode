from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        curr_idx = 0
        zero_idx = len(nums) - 1
        while curr_idx < len(nums) and curr_idx < zero_idx:
            if nums[curr_idx] == 0:
                while zero_idx == 0:
                    zero_idx-=1
                nums[curr_idx] = nums[zero_idx]
                nums[zero_idx] = 0
                zero_idx -= 1
            curr_idx += 1
    
    def moveZeroes2(self, nums: List[int]) -> None:
        lastNonZeroFoundAt = 0
        cur = 0
        while cur < len(nums):
            if nums[cur] != 0:
                nums[lastNonZeroFoundAt], nums[cur] = nums[cur], nums[lastNonZeroFoundAt]
                lastNonZeroFoundAt+=1
            cur+=1

def test_move_zeroes():
    q = [0,1,0,3,12]
    Solution().moveZeroes2(q)
    assert q == [1,3,12,0,0]
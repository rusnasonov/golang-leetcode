from typing import List

class Solution:
    def subarraySumBrutforce(self, nums: List[int], k: int) -> int:
        total_arrays = 0
        for i in range(len(nums)+1):
            for j in range(i+1, len(nums)+1):
                if sum(nums[i:j]) == k:
                    total_arrays+=1
        return total_arrays

    def subarraySumCummulative(self, nums: List[int], k: int) -> int:
        total_arrays = 0
        s = [0]
        for item in nums:
            s.append(s[-1]+item)
        for i in range(len(nums)+1):
            for j in range(i+1, len(nums)+1):
                if s[j] - s[i] == k:
                    total_arrays+=1
        return total_arrays

    def subarraySumCummulativeDict(self, nums: List[int], k: int) -> int:
        pre_sums = {0: 1}
        s = 0
        res = 0
        for num in nums:
            s += num
            res += pre_sums.get(s - k, 0)
            pre_sums[s] = pre_sums.get(s, 0) + 1
        return res
    
    def subarraySum(self, *args, **kwargs) -> int:
        return self.subarraySumCummulative(*args, **kwargs)

def test_subarray_sum():
    assert Solution().subarraySum([1,1,1],2) == 2
    assert Solution().subarraySum([1,2,3],3) == 2
    assert Solution().subarraySum([1],0) == 0
    assert Solution().subarraySum([1,-1,0],0) == 3
    assert Solution().subarraySum([1,-1,0],0) == 3
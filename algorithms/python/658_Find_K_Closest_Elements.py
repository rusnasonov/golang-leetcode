from typing import List
import bisect

class Solution:
    def findClosestElements1(self, arr: List[int], k: int, x: int) -> List[int]:
        # my solution, sliding window, O(n)
        start, end = 0, k - 1
        end_dist = abs(x - arr[end])
        start_dist = abs(x - arr[start])
        while end < len(arr) - 1:
            next_dist = abs(arr[end+1] - x)
            if next_dist >= end_dist and next_dist >= start_dist and arr[end+1] != arr[start]:
                break
            start+=1
            end+=1
            start_dist = abs(x - arr[start])
            end_dist = next_dist
        print(start, end)
        return arr[start:end+1]

    def findClosestElements2(self, arr: List[int], k: int, x: int) -> List[int]:
        # O(log(n-k)) binary search left border
        left, right = 0, len(arr) - k
        while left < right:
            mid = (left + right) // 2
            if x - arr[mid] > arr[mid + k] - x:
                left = mid + 1
            else:
                right = mid
        return arr[left:left + k]

    def findClosestElements3(self, arr: List[int], k: int, x: int) -> List[int]:
        # O(log(n) + k) binary search element (or close element) index, then shrink window
        if x <= arr[0]:
            return arr[:k]
        if x >= arr[-1]:
            return arr[len(arr)-k:len(arr)]
        idx = bisect.bisect(arr, x)
        low = max(0, idx-k-1)
        high = min(len(arr)-1, idx+k-1)
        while high - low > k - 1:
            if low < 0 or (x - arr[low] <= arr[high] - x):
                high-=1
            elif high > len(arr) - 1 or x - arr[low] > arr[high] - x:
                low+=1
        return arr[low:high+1]

    def findClosestElements(self, *args, **kwargs) -> List[int]:
        return self.findClosestElements1(*args, **kwargs)

def test_findClosestElements():
    assert Solution().findClosestElements([1,2,3,4,5], 4, 3) == [1,2,3,4]
    assert Solution().findClosestElements([1,2,3,4,5], 4, -1) == [1,2,3,4]
    assert Solution().findClosestElements([1,2,3,4,10,11,12,18,19,20], 3, 11) == [10,11,12]
    assert Solution().findClosestElements([1,2,3,4,10,11,12,18,19,20], 4, 11) == [4,10,11,12]
    assert Solution().findClosestElements([1,1,1,10,10,10], 1, 9) == [10]
    assert Solution().findClosestElements([1,3,3,4,5,7,7,8,8,8], 6, 6) == [4,5,7,7,8,8]
    assert Solution().findClosestElements([1,1,2,2,2,2,2,3,3], 3, 3) == [2,3,3]
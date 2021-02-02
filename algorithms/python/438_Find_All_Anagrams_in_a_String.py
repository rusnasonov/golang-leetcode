from typing import List
from collections import defaultdict

class Solution:
    def findAnagrams(self, p: str, s: str) -> List[int]:
        if not p or len(p) < len(s):
            return []
        indexes = []
        count = 0
        alphabet_size = 26
        a = ord('a')
        s_freq = [0]*alphabet_size
        p_freq = [0]*alphabet_size
        for i in range(len(s)):
            s_freq[ord(s[i]) - a] += 1
            p_freq[ord(p[i]) - a] += 1
        for i in range(alphabet_size):
            if s_freq[i] == p_freq[i]:
                count+=1
        i = 0
        while i < len(p) - len(s):
            if count == alphabet_size:
                indexes.append(i)
            l = ord(p[i]) - a
            r = ord(p[i + len(s)]) - a
            p_freq[r]+=1
            if p_freq[r] == s_freq[r]:
                count+=1
            elif p_freq[r] == s_freq[r] + 1:
                count-=1
            p_freq[l]-=1
            if p_freq[l] == s_freq[l]:
                count+=1
            elif p_freq[l] == s_freq[l] - 1:
                count-=1
            i+=1
        if count == 26 and len(p) - len(s) == i:
            indexes.append(i)
        return indexes

def test_subarray_sum():
    assert Solution().findAnagrams("cbaebabacd", "abc") == [0, 6]
    assert Solution().findAnagrams("abab", "ab") == [0, 1, 2]
    assert Solution().findAnagrams("aaaaaaaaaa", "aaaaaaaaaaaaa") == []
    assert Solution().findAnagrams("", "a") == []
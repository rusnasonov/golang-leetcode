class Solution:
    def longestPalindrome(self, s: str) -> str:
        # O(n^3) solution
        start_m, end_m = 0, 0
        max_p = 0
        for start in range(len(s)):
            for end in range(len(s)-1, 0, -1):
                if self.is_palindrome(s, start, end) and end - start >= max_p:
                    max_p = end - start
                    start_m = start
                    end_m = end
        return s[start_m:end_m+1]

    def is_palindrome(self, s, start, end):
        l, r = start, end
        while l < r:
            if s[l] != s[r]:
                return False
            r-=1
            l+=1
        return True

    def canonical(self, s):
        # O(n^2), 0(1) space
        start, end = 0, 0
        for i in range(len(s)):
            m1 = self.expand_around_center(s, i, i)
            m2 = self.expand_around_center(s, i, i+1)
            m = max(m1, m2)
            if m > end - start:
                start = i - (m - 1) // 2
                end = i + m // 2
        return s[start:end+1] 

    def expand_around_center(self, s, left, right):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left-=1
            right+=1
        return right - left - 1
    
    def manacher(self, s: str):
        # O(n), not work =(
        if len(s) <=1:
            return s
        s2 = ["#"]*(len(s)*2+1)
        for i in range(len(s)):
            s2[i*2+1] = s[i]
        p = [0]*len(s2)
        r, c, index, i_mir = 0, 0, 0, 0
        max_len = 1
        for i in range(1, len(s2)):
            i_mir = 2*c-i
            p[i] = min(p[i_mir], r-i) if r > i else 0
            while i > p[i] and (i + p[i] + 1) < len(s2) and s2[i - p[i] - 1] == s2[i + p[i] + 1]:
                p[i] += 1
            if p[i] + i > r:
                c = i
                r = i + p[i]
            if max_len < p[i]:
                max_len = p[i]
                index = i
        pos = (index-max_len)//2
        return s[pos:pos+max_len]

def test_longest_palindrome():
    s = Solution()
    assert s.longestPalindrome("babad") == "aba"
    assert s.longestPalindrome("cbbd") == "bb"
    assert s.longestPalindrome("a") == "a"
    assert s.longestPalindrome("ac") == "c"
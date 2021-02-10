from collections import defaultdict, Counter

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not t or not s:
            return ""
        dict_t = Counter(t)
        required = len(dict_t)
        l, r = 0, 0
        formed = 0
        window_counts = defaultdict(int)
        ans = float("inf"), None, None
        filtered_s = [(i, item) for i, item in enumerate(s) if item in dict_t]
        while r < len(filtered_s):
            character = filtered_s[r][1]
            window_counts[character] += 1
            if character in dict_t and window_counts[character] == dict_t[character]:
                formed += 1
            while l <= r and formed == required:
                character = filtered_s[l][1]
                end = filtered_s[r][0]
                start = filtered_s[l][0]
                if end - start + 1 < ans[0]:
                    ans = (end - start + 1, start, end)

                window_counts[character] -= 1
                if window_counts[character] < dict_t[character]:
                    formed -= 1
                l += 1    
            r += 1    
        return "" if ans[0] == float("inf") else s[ans[1] : ans[2] + 1]

def test_min_window():
    s = Solution()
    assert s.minWindow(s="ADOBECODEBANC", t="ABC") == "BANC"
    assert s.minWindow(s="a", t="a") == "a"
    assert s.minWindow(s="a", t="b") == ""
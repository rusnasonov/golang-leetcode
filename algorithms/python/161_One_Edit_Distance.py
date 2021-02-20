class Solution:
    def isOneEditDistance(self, s: str, t: str) -> bool:
        s_len, t_len = len(s), len(t)
        if s_len > t_len:
            return self.isOneEditDistance(t, s)
        if s_len - t_len > 1:
            return False
        for i in range(s_len):
            if s[i] != t[i]:
                if s_len == t_len:
                    return s[i + 1:] == t[i + 1:]
                else:
                    return s[i:] == t[i + 1:]
        return s_len + 1 == t_len

def test_isOneEditDistance():
    assert Solution().isOneEditDistance("cat", "cats") is True
    assert Solution().isOneEditDistance("cat", "catss") is False
    assert Solution().isOneEditDistance("ab", "acb") is True
    assert Solution().isOneEditDistance("", "") is False
    assert Solution().isOneEditDistance("a", "") is True
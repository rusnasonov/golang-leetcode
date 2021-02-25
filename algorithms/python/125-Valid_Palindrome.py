class Solution:
    def isPalindrome(self, s: str) -> bool:
        l, r = 0, len(s)-1
        while l < r:
            l_char = s[l].lower()
            if not l_char.isalnum():
                l+=1
                continue
            r_char = s[r].lower()
            if not r_char.isalnum():
                r-=1
                continue
            if l_char != r_char:
                return False
            r-=1
            l+=1
        return True


def test_is_palindrome():
    s = Solution()
    assert s.isPalindrome("A man, a plan, a canal: Panama") is True 
    assert s.isPalindrome("race a car") is False 
    assert s.isPalindrome("0P") is False 
    assert s.isPalindrome("") is True 
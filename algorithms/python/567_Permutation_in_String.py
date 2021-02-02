from collections import defaultdict

class Solution:
    primes = {'a': 2, 
                  'b': 3, 
                  'c': 5, 
                  'd': 7, 
                  'e': 11, 
                  'f': 13,
                  'g': 17,
                  'h': 19,
                  'i': 23,
                  'j': 29,
                  'k': 31,
                  'l': 37,
                  'm': 41,
                  'n': 43,
                  'o': 47,
                  'p': 53,
                  'q': 59,
                  'r': 61,
                  's': 67, 
                  't': 71,
                  'u': 73,
                  'v': 79,
                  'w': 83,
                  'x': 89,
                  'y': 97,
                  'z': 101
                 }

    def hash(self, s: str) -> int:
        h = 1
        for c in s:
            h *= self.primes[c]
        return h

    def checkInclusion(self, s1: str, s2: str) -> bool:
        h1 = self.hash(s1)
        l = len(s1)
        h2 = self.hash(s2[:l])
        if h1 == h2:
            return True
        for i in range(1, len(s2)-l+1):
            h2 //= self.primes[s2[i-1]]
            h2 *= self.primes[s2[i+l-1]]
            if h1 == h2:
                print(i)
                return True
        return False
    
    def checkInclusionDict(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
        s1_dict = defaultdict(int)
        s2_dict = defaultdict(int)
        count = 0
        alphabet_size = 26
        for i in range(len(s1)):
            s1_dict[ord(s1[i])-ord('a')] += 1
            s2_dict[ord(s2[i])-ord('a')] += 1
        for i in range(alphabet_size):
            if s1_dict[i] == s2_dict[i]:
                count+=1
        for i in range(len(s2) - len(s1)):
            if count == 26:
                return True
            r = ord(s2[i + len(s1)]) - ord('a')
            l = ord(s2[i]) - ord('a')
            s2_dict[r]+=1
            if s2_dict[r] == s1_dict[r]:
                count+=1
            elif s2_dict[r] == s1_dict[r] + 1:
                count-=1
            s2_dict[l]-=1
            if s2_dict[l] == s1_dict[l]:
                count+=1
            elif s2_dict[l] == s1_dict[l] - 1:
                count-=1
        return count == alphabet_size

def test_subarray_sum():
    assert Solution().checkInclusion("ab", "eidbaooo") == True
    assert Solution().checkInclusion("ab", "eidboaoo") == False
    assert Solution().checkInclusion("adc", "dcda") == True
    assert Solution().checkInclusion("trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine") == True

    assert Solution().checkInclusionDict("ab", "eidbaooo") == True
    assert Solution().checkInclusionDict("ab", "eidboaoo") == False
    assert Solution().checkInclusionDict("adc", "dcda") == True
    assert Solution().checkInclusionDict("trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine") == True
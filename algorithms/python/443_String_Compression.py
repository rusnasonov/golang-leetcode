from typing import List

class Solution:
    def compress(self, chars: List[str]) -> int:
        if not chars:
            return []
        cur, next, write_to, cnt = 0, 1, 0, 1
        while next <= len(chars):
            if next < len(chars) and chars[cur] == chars[next]:
                cnt+=1
            else:
                chars[write_to] = chars[cur]
                write_to += 1
                cur = next
                if cnt > 1:
                    cnt_s = str(cnt)
                    for i in cnt_s:
                        chars[write_to] = i
                        write_to += 1
                    cnt = 1
            next+=1
        return write_to#list("".join(chars[:write_to]))
    
    def canonical(self, chars):
        anchor = write = 0
        for read, c in enumerate(chars):
            if read + 1 == len(chars) or chars[read + 1] != c:
                chars[write] = chars[anchor]
                write += 1
                if read > anchor:
                    for digit in str(read - anchor + 1):
                        chars[write] = digit
                        write += 1
                anchor = read + 1
        return write

def test_compress():
    s = Solution()
    assert s.compress(["a","a","b","b","c","c","c"]) == 6#["a","2","b","2","c","3"]
    assert s.compress(["a"]) == 1#["a"]
    assert s.compress(["a","b","b","b","b","b","b","b","b","b","b","b","b"]) == 4#["a","b","1","2"]
    assert s.compress(["a","a","a","b","b","a","a"]) == 6#["a","3","b","2","a","2"]
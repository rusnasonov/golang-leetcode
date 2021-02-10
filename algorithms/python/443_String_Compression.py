from typing import List

class Solution:
    def compress(self, chars: List[str]) -> int:
        if not chars:
            return []
        cur, next, write_to, cnt = 0, 1, 0, 1
        while cur < next:
            if chars[cur] == chars[next]:
                cnt+=1
            else:
                chars[write_to] = chars[cur]
                write_to += 1
                cur = next
                if cnt > 1:
                    chars[write_to] = str(cnt)
                    write_to += 1
                    cnt = 1
            next+=1
        return chars[:write_to]


def test_compress():
    s = Solution()
    assert s.compress(["a","a","b","b","c","c","c"]) == ["a","2","b","2","c","3"]
    assert s.compress(["a"]) == ["a"]
    assert s.compress(["a","b","b","b","b","b","b","b","b","b","b","b","b"]) == ["a","b","1","2"]
    assert s.compress(["a","a","a","b","b","a","a"]) == ["a","3","b","2","a","2"]
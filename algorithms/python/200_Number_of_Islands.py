from typing import List

class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        rc = len(grid)
        cc = len(grid[0])
        cells = dict()
        for i in range(rc):
            for j in range(cc):
                if grid[i][j] == '1':
                    cells[(i,j)] = 0 
        queue = set()
        islands = 0
        while len(cells) > 0:
            cell = next(iter(cells))
            del cells[cell]
            if grid[cell[0]][cell[1]] == '1':
                queue.add(cell)
                islands += 1
                while len(queue) > 0:
                    item = queue.pop()
                    r, c = item
                    bottom = (r + 1, c) if r + 1 < rc else None
                    top = (r - 1, c) if r - 1 < rc else None
                    right = (r, c + 1) if c + 1 < cc else None
                    left = (r, c - 1) if c - 1 < cc else None
                    for c in (bottom, top, right, left):
                        if c and grid[c[0]][c[1]] == '1' and c in cells:
                            queue.add(c)
                            del cells[c]
        return islands

def test_num_islands():
    grid = [
        ["1","1","1","1","0"],
        ["1","1","0","1","0"],
        ["1","1","0","0","0"],
        ["0","0","0","0","0"]
    ]
    out = 1
    assert Solution().numIslands(grid) == out
    grid = [
        ["1","1","0","0","0"],
        ["1","1","0","0","0"],
        ["0","0","1","0","0"],
        ["0","0","0","1","1"]
    ]
    out = 3
    assert Solution().numIslands(grid) == out
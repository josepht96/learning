class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        Point = namedtuple('point', ['r', 'c'])
        seen = []
        queue = []
        biggest_island = 0
        for r, row in enumerate(grid):
            for c, column in enumerate(row):
                if Point(r, c) in seen:
                    continue
                # if r,c is an island, perform BFS
                # add to queue, then add to seen
                if grid[r][c] == 1:
                    queue.append(Point(r, c))
                    current_island_size = 0
                    while len(queue) != 0:
                        point = queue.pop()
                        if point in seen:
                            continue
                        current_island_size += 1
                        # check left
                        if point.c != 0:
                            if grid[point.r][point.c-1] == 1:
                                queue.append(Point(point.r, point.c-1))
                        # check right
                        if point.c != len(grid[0])-1:
                            if grid[point.r][point.c+1] == 1:
                                queue.append(Point(point.r, point.c+1))
                        # check up
                        if point.r != 0:
                            if grid[point.r-1][point.c] == 1:
                                queue.append(Point(point.r-1, point.c))
                        # check down
                        if point.r != len(grid)-1:
                            if grid[point.r+1][point.c] == 1:
                                queue.append(Point(point.r+1, point.c))
                        seen.append(Point(point.r, point.c))
                
                    if current_island_size > biggest_island:
                        biggest_island = current_island_size
        return biggest_island
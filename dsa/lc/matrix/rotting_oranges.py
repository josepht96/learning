class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        # add rotten to queue
        time_grid = deepcopy(grid)
        rotten_queue = []
        fresh_queue = []
        for r in range(len(grid)):
            for c in range(len(grid[0])):
                if grid[r][c] == 2:
                    rotten_queue.append([r, c])
                    time_grid[r][c] = 0
                elif grid[r][c] == 1:
                    fresh_queue.append([r, c])
        max = 0
        while rotten_queue:
            r, c = rotten_queue.pop(0)
            if [r, c] in fresh_queue:
                fresh_queue.remove([r, c])
            if time_grid[r][c] > max:
                max = time_grid[r][c]
            # check left
            if c != 0:
                if grid[r][c-1] == 1:
                    time_grid[r][c-1] = time_grid[r][c] + 1
                    rotten_queue.append([r, c-1])
                    grid[r][c-1] = 2

            # check right
            if c != len(grid[0]) - 1:
                if grid[r][c+1] == 1:
                    time_grid[r][c+1] = time_grid[r][c] + 1
                    rotten_queue.append([r, c+1])
                    grid[r][c+1] = 2

            # check up
            if r != 0:
                if grid[r-1][c] == 1:
                    time_grid[r-1][c] = time_grid[r][c] + 1
                    rotten_queue.append([r-1, c])
                    grid[r-1][c] = 2

            # check down
            if r != len(grid)-1:
                if grid[r+1][c] == 1:
                    time_grid[r+1][c] = time_grid[r][c] + 1
                    rotten_queue.append([r+1, c])
                    grid[r+1][c] = 2

        if fresh_queue:
            return -1

        return max
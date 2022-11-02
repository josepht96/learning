class Solution:
    def findBall(self, grid: List[List[int]]) -> List[int]:
        # i == row
        # j == column
        returnList = []
        for ballNum, L in enumerate(grid[0]):
            outcome = self.dropBall(grid, 0, ballNum)
            returnList.append(outcome)
            
        return returnList
        
    def dropBall(self, grid, row: int, col: int) -> int:
        # going left
        if grid[row][col] == -1:
            if col == 0 or grid[row][col-1] == 1:
                return -1
            elif row == len(grid)-1:
                return col-1
            else:
                return self.dropBall(grid, row+1, col-1)
            
        # going right
        elif grid[row][col] == 1:
            if col == len(grid[0])-1 or grid[row][col+1] == -1:
                return -1
            elif row == len(grid)-1:
                return col+1
            else:
                return self.dropBall(grid, row+1, col+1)
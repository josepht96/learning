# bfs for entire data set or bfs for each point

from copy import deepcopy
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        returnMat = deepcopy(mat)
        queue = []
        for r in range(len(mat)):
            for c in range(len(mat[0])):
                if mat[r][c] == 0:
                    queue.append([r, c])
                    returnMat[r][c] = 0
                else:
                    returnMat[r][c] = 9999999999999
        seen = []
        while queue:
            r, c = queue.pop(0)
            dist = returnMat[r][c]
            # check left
            if c != 0:
                if dist < returnMat[r][c-1]:
                    returnMat[r][c-1] = dist + 1
                    queue.append([r, c-1])
            # check right
            if c != len(mat[0])-1:
                if dist < returnMat[r][c+1]:
                    returnMat[r][c+1] = dist + 1
                    queue.append([r, c+1])
            # check up
            if r != 0:
                if dist < returnMat[r-1][c]:
                    returnMat[r-1][c] = dist + 1
                    queue.append([r-1, c])
            # check down
            if r != len(mat)-1:
                if dist < returnMat[r+1][c]:
                    returnMat[r+1][c] = dist + 1
                    queue.append([r+1, c])

        return returnMat
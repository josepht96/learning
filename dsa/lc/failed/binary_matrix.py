class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        Point = namedtuple('point', ['r', 'c'])
        returnMat = mat.copy()
        # if zero do nothing
        # if not zero, perform BFS
        checked = []
        for r in range(len(mat)):
            for c in range(len(mat[0])):
                if mat[r][c] != 0:
                    queue = []
                    seen = []
                    queue.append(Point(r, c))
                    while queue:
                        p = queue.pop(0)
                        if mat[p.r][p.c] == 0:
                            # calculate distance to r, c
                            horiz_dist = abs(c - p.c)
                            vert_dist = abs(r - p.r)
                            returnMat[r][c] = horiz_dist + vert_dist
                            break
                        if p.c != 0:
                            if Point(p.r, p.c-1) not in seen:
                                queue.append(Point(p.r, p.c-1))
                        if p.c != len(mat[0])-1:
                            if Point(p.r, p.c+1) not in seen:
                                queue.append(Point(p.r, p.c+1))
                        if p.r != 0:
                            if Point(p.r-1, p.c) not in seen:
                                queue.append(Point(p.r-1, p.c))
                        if p.r != len(mat)-1:
                            if Point(p.r+1, p.c) not in seen:
                                queue.append(Point(p.r+1, p.c))
                        seen.append(Point(p.r, p.c))
        return returnMat

from copy import deepcopy
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        Point = namedtuple('point', ['r', 'c'])
        distMat = deepcopy(mat)
        # if zero do nothing
        # if not zero, perform BFS
        checked = []
        for r in range(len(mat)):
            for c in range(len(mat[0])):
                if mat[r][c] == 0:
                    queue = []
                    seen = []
                    queue.append(Point(r, c))
                    while queue:
                        p = queue.pop(0)
                        if mat[p.r][p.c] == 1:
                            hdist = abs(c - p.c)
                            vdist = abs(r - p.r)
                            dist = hdist + vdist
                            # determine if current point is closer than current distance
                            if Point(p.r, p.c) in checked and dist < returnMat[p.r][p.c]:
                                returnMat[p.r][p.c] = dist
                            elif Point(p.r, p.c) not in checked:
                                checked.append(Point(p.r, p.c))
                                returnMat[p.r][p.c] = dist

                        if p.c != 0 and mat[p.r][p.c-1] == 1:
                            if Point(p.r, p.c-1) not in seen:
                                queue.append(Point(p.r, p.c-1))
                                
                        if p.c != len(mat[0])-1 and mat[p.r][p.c+1] == 1:
                            if Point(p.r, p.c+1) not in seen:
                                queue.append(Point(p.r, p.c+1))

                        if p.r != 0 and mat[p.r-1][p.c] == 1:
                            if Point(p.r-1, p.c) not in seen:
                                queue.append(Point(p.r-1, p.c))

                        if p.r != len(mat)-1 and mat[p.r+1][p.c] == 1:
                            if Point(p.r+1, p.c) not in seen:
                                queue.append(Point(p.r+1, p.c))

                        seen.append(Point(p.r, p.c)) 
        return returnMat
                            
                        
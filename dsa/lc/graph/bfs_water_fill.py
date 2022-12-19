class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        # start at List[sr][sc]
        # change color to new color
        # any spots connected on a that face get the color
        # any spots connected to those spots get the color, etc...
        # Start at start_point, add each adjacent spot to queue
        # check each adjacent, then check their adjacents, etc...
        
        Point = namedtuple('point', ['r', 'c'])
        queue = []
        seen = []
        original_color = image[sr][sc]
        queue.append(Point(sr, sc))
        #process each point
        #check adjacent points
        while len(queue) != 0:
            point = queue.pop()
            if point in seen:
                continue
            image[point.r][point.c] = color
            # process adjacent nodes
            # check left
            if point.c != 0:
                if image[point.r][point.c-1] == original_color:
                    queue.append(Point(point.r, point.c-1))

            # check right
            if point.c != len(image[0])-1:
                if image[point.r][point.c+1] == original_color:
                    queue.append(Point(point.r, point.c+1))

            # check up
            if point.r != 0:
                if image[point.r-1][point.c] == original_color:
                    queue.append(Point(point.r-1, point.c))

            # check down
            if point.r != len(image)-1:
                if image[point.r+1][point.c] == original_color:
                    queue.append(Point(point.r+1, point.c))

            seen.append(Point(point.r, point.c))
        return image
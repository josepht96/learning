func min(x, y int) int {
    if x > y {
        return y
    } else {
        return x
    }
}

func maxArea(height []int) int {
    leftIndex := 0
    rightIndex := len(height) - 1
    maxArea := 0
    for leftIndex < rightIndex {
        l := rightIndex - leftIndex
        minH := min(height[leftIndex], height[rightIndex])
        area := l * minH
        if area > maxArea {
            maxArea = area
        }
        if height[leftIndex] > height[rightIndex] {
            rightIndex--
        } else {
            leftIndex++
        }
    }
    return maxArea
    
}
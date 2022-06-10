func min(x, y int) int {
    if x >= y {
        return y
    } else {
        return x
    }
}

func maxArea(height []int) int {
    maxArea := 0
    for i:=1; i<len(height); i++ {
        for j:=0; j<i; j++ {
            l := i - j
            h := min(height[j], height[i])
            area := l * h
            if area > maxArea {
                maxArea = area
            }
        }
    }
    return maxArea
}
func transpose(matrix [][]int) [][]int {
    height := len(matrix[0])
    width := len(matrix)
    var newMatrix [][]int
    for i:=0; i<height;i++ {
        var newArray []int
        for j:=0; j<width;j++ {
            newArray = append(newArray, matrix[j][i])
        }
        newMatrix = append(newMatrix, newArray)
    }
    return newMatrix
}
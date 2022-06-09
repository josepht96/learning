func twoSum(numbers []int, target int) []int {
    m := make(map[int]int)
    for i:=0; i<len(numbers); i++ {
        diff := target - numbers[i]
        value, ok := m[numbers[i]]
        if ok {
            resp := []int{value + 1, i + 1}
            return resp
        } else {
            m[diff] = i
        }
    }
    return nil
}
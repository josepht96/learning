func twoSum(nums []int, target int) []int {
    //take first number
    //determine remainder for each spot
    //if succeeding number matches remainder, you are done
    m := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        //determine remainder
        remainder := target - nums[i]
        //check if remainder exists in map
        value, ok := m[remainder]
        //if value is present return array
        if ok {
            x := []int{value, i}
            return x
        } else {
            m[nums[i]] = i
        } 
    }
    return nil
}
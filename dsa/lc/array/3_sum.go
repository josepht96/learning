func threeSum(nums []int) [][]int {
    //two pointers
    //sort array
    //start at beginning and end, sum them
    //if sum is less than zero, move left point
    //else move right pointer
    //if start pointer == prev, skip
    
    //sort array
    var returnArr [][]int
    sort.Ints(nums[:])
    
    for i:=0; i < len(nums) && nums[i] <= 0; i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        
        lo := i+1
        high := len(nums) - 1
        for lo < high {
            sum := nums[i] + nums[lo] + nums[high]
            if sum == 0 {
                tmpArr := []int{nums[i], nums[lo], nums[high]}
                returnArr = append(returnArr, tmpArr)
                high--
                for lo < high && nums[lo] == nums[lo+1] {
                    lo++
                }
                lo++
            } else if sum < 0 {
                lo++
            } else {
                high--
            }
        }       
    }
    return returnArr
}
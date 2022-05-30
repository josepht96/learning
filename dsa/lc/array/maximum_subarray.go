func maxSubArray(nums []int) int {
    // O(n) is possible which means you need to track trailing probably
    //tracking current plus previous sum gives you current sum
    //if trailing value is less than current value probably can discard it
    //find the starting high point
    //end point would then just be the highest value

    //can do it with array but appending is much slower than altering a single value
    highestValue := nums[0]
    prefixValue := nums[0]
    for i:=1; i<len(nums);i++ {
        if (prefixValue+nums[i] > nums[i]) {
            //append to prefix path but not certain its better
            prefixValue = prefixValue + nums[i]
        } else {
            //change  prefixArr value for next iteration
            prefixValue = nums[i]
        }

        if prefixValue > highestValue {
            highestValue = prefixValue
        }
    }
    return highestValue
}
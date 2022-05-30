func maxProduct(nums []int) int {
    //make sure highestProduct starts with nums[0]
    highestProduct := nums[0]
    prefixProduct := nums[0]
    prefixProductAll := nums[0]
    zeroExists := false
    for i:=1; i<len(nums); i++ {
        //track negative and positive
        //if current value is higher than prefixProduct, 
        //negative values are good if there is an even amount
        //zero values need to be handled
        //zero value should reset both pos and negative
        //tracking
        //prefixProduct gets current value
        //if current prefix is best then set highestProduct
        
        //negative value is always incremented
        //handle zero
        if nums[i] == 0 {
            prefixProductAll = 1
            prefixProduct = 1
            zeroExists = true
            continue
        }
        
        //this should always be done
        prefixProductAll = prefixProductAll * nums[i]
        tempProduct := prefixProduct * nums[i]
        
        if nums[i] > tempProduct {
            prefixProduct = nums[i]
        } else {
            prefixProduct = tempProduct
        }
        
        if prefixProductAll > prefixProduct {
            prefixProduct = prefixProductAll
        }
        
        if prefixProduct > highestProduct {
            highestProduct = prefixProduct
        }
    }
    if highestProduct < 0 && zeroExists {
        return 0
    }
    return highestProduct
}
func productExceptSelf(nums []int) []int {
    //two arrays
    //traverse forward and record the prefix product in answer array
    //starts at 1
    //then multiple the prefix product by the postfix product
    //multiple the current postfix value by the previous postfix product
    
    var prefixArr []int
    var currentPostfixProduct = 1
    var currentPrefix int
    //initialize array 
    prefixArr = append(prefixArr, 1)
    //fill out prefix array
    for i:=1; i<len(nums); i++ {
		//product of the previous num times the prefixArr value of the previous index
        currentPrefix = prefixArr[i-1] * nums[i-1]
        prefixArr = append(prefixArr, currentPrefix)
    }
    //traverse backward
    for i := len(nums) - 2; i >= 0; i-- {
        currentPostfixProduct = currentPostfixProduct * nums[i+1]
        prefixArr[i] = prefixArr[i] * currentPostfixProduct
    }
    
    return prefixArr
}
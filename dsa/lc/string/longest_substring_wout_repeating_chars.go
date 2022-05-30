func lengthOfLongestSubstring(s string) int {
    //iterate through string
    //track chars in map
    //need to track which characters are still relevant
    
    if len(s) == 0 {
        return 0
    }
    maxLength := 1
    currentLength := 1
    m := make(map[string]bool)
    m[string(s[0])] = true
    for i:=1; i<len(s); i++ {
        if m[string(s[i])] != true {
            m[string(s[i])] = true
            currentLength++
        } else if m[string(s[i])] == true {
            //clear map
            m = make(map[string]bool)
            m[string(s[i])] = true
            currentLength = 1
            j := i - 1
            for m[string(s[j])] != true{
                m[string(s[j])] = true
                currentLength++
                j--
            }
        }
        
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }
    return maxLength
}
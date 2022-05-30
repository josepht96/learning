func strStr(haystack string, needle string) int {
    //take first char, check if starting there matches
    //iterate
    //feels like brute force a bit
    lenH:=len(haystack)
    lenN:=len(needle)
    
    if(lenN==0) {
        return 0
    }
    //only need to check until theres enough remaining chars for the needle to match
    for i:=0; i<=lenH-lenN; i++ {
        j:=0
        for j<lenN && haystack[i+j]==needle[j] {
            j++;
        }
        if j==lenN {
            return i;
        }
    }
    return -1;
}
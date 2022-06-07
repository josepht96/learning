func merge(nums1 []int, m int, nums2 []int, n int)  {
    nSpot := 0
    var newArr []int
    for i:=0; i<len(nums1); i++ {
        for j:=nSpot; j<len(nums2); j++ {
            if nums1[i] <= nums2[j] && nums1[i] != 0 {
                newArr = append(newArr, nums1[i])
                break
            } else {
                newArr = append(newArr, nums2[j])
                nSpot++
            }
        }
    }
    nums1 = newArr
}
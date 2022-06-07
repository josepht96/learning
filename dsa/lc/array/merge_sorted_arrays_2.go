func merge(nums1 []int, m int, nums2 []int, n int)  {
    j, k := 0, 0
    nums1Copy := make([]int, len(nums1))
    copy(nums1Copy, nums1)
    //w1 = &nums1[i]
    for i:=0; i<len(nums1); i++ {
        if k == n {
        	nums1[i] = nums1Copy[j]
        	j++
            continue
        }
        if nums1Copy[j] <= nums2[k] && j < m {
            nums1[i] = nums1Copy[j]
            j++
        } else {
            nums1[i] = nums2[k]
            k++
        }
    }
}
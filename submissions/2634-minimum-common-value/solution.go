func getCommon(nums1 []int, nums2 []int) int {
    l := 0; r := 0

    for l < len(nums1) && r < len(nums2) {
        if nums1[l] == nums2[r] {
            return nums1[l]
        }
        
        if nums1[l] < nums2[r] {
            for l < len(nums1) && nums1[l] < nums2[r] {
                l++
            }
            
            continue
        }

        if nums1[l] > nums2[r] {
            for r < len(nums2) && nums1[l] > nums2[r] {
                r++
            }

            continue
        }
    }

    return -1
}

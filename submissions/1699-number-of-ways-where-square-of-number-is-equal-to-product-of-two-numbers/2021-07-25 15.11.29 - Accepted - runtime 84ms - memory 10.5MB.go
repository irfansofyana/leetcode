func numTriplets(nums1 []int, nums2 []int) int {
    n := []int{len(nums1), len(nums2)}
    arr := [][]int{nums1, nums2}
    ans := 0
    
    for i := 0; i < 2; i++ {
        isExist := map[int64]int{}
        for j := 0; j < n[i]-1; j++ {
            for k := j + 1; k < n[i]; k++ {
                if _, ok := isExist[(int64)(arr[i][j]*arr[i][k])]; ok {
                    isExist[(int64)(arr[i][j]*arr[i][k])] += 1
                } else {
                    isExist[(int64)(arr[i][j]*arr[i][k])] = 1
                }
            }
        }
        for j := 0; j < n[1-i]; j++ {
            if val, ok := isExist[(int64)(arr[1-i][j]*arr[1-i][j])]; ok {
                ans += val
            }
        }
    }
    
    return ans
}
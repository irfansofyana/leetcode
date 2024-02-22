func subsets(nums []int) [][]int {
    subset := make([][]int, 0); n := len(nums)
    
    for i := 0; i < (1 << n); i++ {
        tmp := make([]int, 0)
        for j := 0; j < n; j++ {
            if (i&(1<<j)) > 0 {
                tmp = append(tmp, nums[j])
            }
        }
        subset = append(subset, tmp)
    }
    return subset
}
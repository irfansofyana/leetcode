func findPrefixScore(nums []int) []int64 {
    n := len(nums)
    conv := make([]int64, n)
    maks := -1
    for i := 0; i < n; i++ {
        if nums[i] > maks {
            maks = nums[i]
        }

        conv[i] = int64(nums[i]) + int64(maks)
        if i > 0 {
            conv[i] += conv[i-1]
        }
    }

    return conv
}

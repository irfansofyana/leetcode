import "sort"

func minPairSum(nums []int) int {
    sort.Ints(nums)
    l := 0; r := len(nums) - 1
    maks := 0
    
    for l <= r {
        if nums[l] + nums[r] > maks {
            maks = nums[l] + nums[r]
        }
        l += 1
        r -= 1
    }
    
    return maks
}

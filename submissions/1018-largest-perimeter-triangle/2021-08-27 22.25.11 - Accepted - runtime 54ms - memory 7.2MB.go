import "sort"

func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    
    res := 0
    for i := 2; i < len(nums); i++ {
        sumTwoSmallestSide := nums[i-1] + nums[i-2]
        if nums[i] < sumTwoSmallestSide {
            if nums[i] + sumTwoSmallestSide > res {
                res = nums[i] + sumTwoSmallestSide
            }
        }
    }    
    
    return res
}
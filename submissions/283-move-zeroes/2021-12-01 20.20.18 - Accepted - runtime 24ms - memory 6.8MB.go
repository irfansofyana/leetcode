func moveZeroes(nums []int)  {
    zeros := 0
    n := len(nums)
    
    for i, num := range nums {
        if num == 0 {
            zeros++
            continue
        }
        nums[i-zeros] = num
    }
    for i := n-zeros; i < n; i++ {
        nums[i] = 0
    }
}
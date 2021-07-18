func removeElement(nums []int, val int) int {
    var result int = 0
    
    for _, num := range(nums) {
        if num != val {
            nums[result] = num
            result += 1
        }
    }
    
    return result
}

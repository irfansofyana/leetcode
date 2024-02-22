func removeDuplicates(nums []int) int {
    var result = 0
    var isExist = map[int]bool{}
    
    for _, num := range(nums) {
        if !isExist[num] {
            result += 1
            nums[result-1] = num
            isExist[num] = true
        }
    }
    
    return result
}
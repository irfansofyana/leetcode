func findMaxConsecutiveOnes(nums []int) int {
    maxOne := 0
    
    i := 0
    for i < len(nums) {
        if nums[i] == 0 {
            i++
            continue
        }
        
        j := i + 1
        freq := 1
        for j < len(nums) && nums[j] == nums[i] {
            freq++
            j++
        }
        
        if freq > maxOne {
            maxOne = freq
        }
        
        i = j
    }
    
    return maxOne
}

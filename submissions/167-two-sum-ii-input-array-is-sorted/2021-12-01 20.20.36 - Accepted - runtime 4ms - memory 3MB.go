func twoSum(numbers []int, target int) []int {
    left := 0; right := 1; n := len(numbers)
    
    ans := make([]int, 2)
    for left < n {
        for right >= 0 && numbers[left] + numbers[right] > target {
            right--
        }
        for right < n && numbers[right] + numbers[left] < target {
            right++
        }
        if right < n && numbers[left] + numbers[right] == target {
            ans[0] = left + 1; ans[1] = right + 1
            break
        }
        left++
        if right == n {
            right--
        }
    }
    
    return ans
}
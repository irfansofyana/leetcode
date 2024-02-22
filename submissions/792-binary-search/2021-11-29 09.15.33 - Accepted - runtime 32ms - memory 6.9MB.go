func search(nums []int, target int) int {
    n := len(nums)
    l := 0; r := n-1
    ans := -1
    
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] == target {
            ans = mid
            break
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    
    return ans
}
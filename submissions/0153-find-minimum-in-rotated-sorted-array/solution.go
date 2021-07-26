func findMin(nums []int) int {
    n := len(nums)
    l := 0; r := n-1
    ans := nums[0]
    
    if n == 1 || nums[0] < nums[n-1] {
        return ans
    }
    
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] > nums[mid+1] {
            ans = nums[mid+1]
            break
        } else if nums[mid-1] > nums[mid] {
            ans = nums[mid]
            break
        } else if nums[mid] > nums[0] {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    
    return ans
}

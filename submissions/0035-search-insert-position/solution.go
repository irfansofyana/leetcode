func searchInsert(nums []int, target int) int {
    var low int = 0
    var high int = len(nums)-1
    var last int = len(nums)
    
    for low <= high {
        var mid int = (low + high) / 2
        
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            low = mid + 1
        } else if nums[mid] > target {
            high = mid - 1
            last = mid
        }
    }
    
    return last
}

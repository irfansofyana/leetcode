func productExceptSelf(nums []int) []int {
    var l = make([]int, 0)
    var r = make([]int, 0)
    var ans = make([]int, 0)
    n := len(nums)
    
    for i := 0; i < n; i++ {
        if i == 0 {
            l = append(l, nums[i])
            r = append(r, nums[n-1-i])
        } else {
            l = append(l, nums[i] * l[i-1])
            r = append(r, nums[n-i-1] * r[i-1])
        }
    }
    
    for i := 0; i < n; i++ {
        if i == 0 {
            ans = append(ans, r[n-2])
        } else if i == n-1 { 
            ans = append(ans, l[n-2])
        } else {
            ans = append(ans, l[i-1] * r[n-2-i])
        }
    }
    
    return ans
}
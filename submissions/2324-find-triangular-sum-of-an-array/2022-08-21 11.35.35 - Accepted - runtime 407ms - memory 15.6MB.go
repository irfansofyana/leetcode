func triangularSum(nums []int) int {
    n := len(nums)
    pascal := make([][]int, n)
    
    for i := 0; i < n; i++ {
        tmp := make([]int, 0)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                tmp = append(tmp, 1)
            } else {
                tmp = append(tmp, (pascal[i-1][j] + pascal[i-1][j-1]) % 10)
            }
        }
        pascal[i] = tmp 
    }
    
    ans := 0
    for i := 0; i < n; i++ {
        ans = (ans + (pascal[n-1][i] * nums[i]) % 10) % 10
    }
    
    return ans
}
func rev(num int) int {
    ret := 0
    for num > 0 {
        ret = 10 * ret + num % 10
        num /= 10
    }
    return ret
}

func countNicePairs(nums []int) int {
    ans := 0
    freq := map[int]int{}
    n := len(nums)
    MOD := 1000000007
    
    for i := 0; i < n; i++ {
        freq[nums[i] - rev(nums[i])] += 1
    }
    
    for _, v := range freq {
        var tmp int64 = int64(v) * int64(v-1) / 2
        tmp %= int64(MOD)
        ans = (ans + int(tmp)) % MOD
    }
    return ans
}

// num[i] - rev(num[i]) = num[j] - rev(num[j])
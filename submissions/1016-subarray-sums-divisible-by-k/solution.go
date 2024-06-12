func subarraysDivByK(nums []int, k int) int {
    // i, j
    // prefix[j] - prefix[i-1] == 0 mod k
    // prefix[j] = prefix[i-1]
    n := len(nums)
    prefix := make([]int, n+1)
    cnt := make([]int, k)

    for i := 0; i < n; i++ {
        prefix[i+1] += prefix[i] + nums[i]
        cnt[findMod(prefix[i+1], k)]++
    }

    ans := 0
    for i := 0; i < k; i++ {
        if i == 0 {
            ans += (cnt[i] * (cnt[i] + 1)) / 2
        } else {
            ans += (cnt[i] * (cnt[i] - 1)) / 2
        }
    }

    return ans
}

func findMod(a, b int) int {
    if a >= 0 {
        return a % b
    }

    return (a + ((-1 * a) / b + 1) * b) % b 
}

// 4 4 4 2 4 0

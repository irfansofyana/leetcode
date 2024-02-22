func sumArr(arr []int) int {
    n := len(arr); sum := 0
    for i := 0; i < n; i++ {
        sum += arr[i]
    }
    return sum
}

func sumOddLengthSubarrays(arr []int) int {
    ans := 0; n := len(arr)
    for i := 0; i < n; i++ {
        for j := 1; i + j - 1 < n; j += 2 {
            ans += sumArr(arr[i:i+j])
        }
    }
    return ans
}
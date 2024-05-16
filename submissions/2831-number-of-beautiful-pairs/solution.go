func countBeautifulPairs(nums []int) int {
    cnt := 0; n := len(nums)

    for i := 0; i < n - 1; i++ {
        for j := i + 1; j < n; j++ {
            firstDig := getFirstDigit(nums[i])
            lastDig := getLastDigit(nums[j])
            if (gcd(firstDig, lastDig) == 1) {
                cnt++
            }
        }
    }

    return cnt
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }

    return gcd(b, a % b)
}

func getFirstDigit(num int) int {
    var dig int = 0
    for num > 0 {
        dig = num % 10
        num /= 10
    }

    return dig
}

func getLastDigit(num int) int {
    return num % 10
}

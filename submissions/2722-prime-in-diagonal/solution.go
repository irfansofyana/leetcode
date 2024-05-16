func diagonalPrime(nums [][]int) int {
    ans := 0

    for i := 0; i < len(nums); i++ {
        num1 := nums[i][i]
        num2 := nums[i][len(nums)-1-i]
        if isPrime(num1) && num1 > ans {
            ans = num1
        }
        if (isPrime(num2) && num2 > ans) {
            ans = num2
        }
    }

    return ans
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }

    if n == 2 {
        return true
    }

    for i := 2; int64(i) * int64(i) <= int64(n); i++ {
        if n % i == 0 {
            return false
        }
    }

    return true
}

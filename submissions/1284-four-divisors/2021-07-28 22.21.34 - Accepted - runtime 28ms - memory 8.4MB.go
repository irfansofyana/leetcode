func isCubicNumber(num int) int {
    l := 0; r := 1000
    cubeRoot := -1
    for l <= r {
        mid := (l + r) / 2
        if mid * mid * mid == num {
            cubeRoot = mid
            break
        } else if mid * mid * mid < num {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return cubeRoot
}

func sumFourDivisors(nums []int) int {
    sieve := make([]int, 100001)
    minFactor := make([]int, 100001)
    
    for i := 2; i <= 100000; i++ {
        if sieve[i] != 0 {
            continue
        }
        for j := i; j <= 100000; j += i {
            sieve[j] += 1
            if minFactor[j] == 0 {
                minFactor[j] = i
            }
        }
        sieve[i] = 0
    }
    
    n := len(nums)
    ans := 0
    for i := 0; i < n; i++ {
        if nums[i] == 1 {
            continue
        }
        if sieve[nums[i]] == 2 {
            minFac1 := minFactor[nums[i]]; num := nums[i] / minFac1
            if num % minFac1 == 0 {
                continue
            }
            minFac2 := minFactor[num]; num /= minFac2
            if num % minFac2 == 0 {
                continue
            }
            ans += (1 + minFac1) * (1 + minFac2)
        } else if isCubicNumber(nums[i]) != -1 {
            cubeRoot := isCubicNumber(nums[i])
            if sieve[cubeRoot] != 0 {
                continue
            }
            ans += (1 + cubeRoot + cubeRoot * cubeRoot + cubeRoot * cubeRoot * cubeRoot)
        }
    }
    return ans
}
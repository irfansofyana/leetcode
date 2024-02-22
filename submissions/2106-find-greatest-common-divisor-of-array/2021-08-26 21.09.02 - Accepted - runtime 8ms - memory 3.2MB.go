func findMin(nums []int) int {
    var min int
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            min = nums[i]
            continue
        }
        if nums[i] < min {
            min = nums[i]
        }
    }
    return min
}

func findMax(nums []int) int {
    var maks int
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            maks = nums[i]
            continue
        }
        if nums[i] > maks {
            maks = nums[i]
        }
    }
    return maks
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func findGCD(nums []int) int {
    min := findMin(nums)
    max := findMax(nums)
    return gcd(min, max)
}
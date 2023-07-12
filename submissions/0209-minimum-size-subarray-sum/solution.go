import "math"

func minSubArrayLen(target int, nums []int) int {
    minimum := math.MaxInt64
    left := 0; right := 0; n := len(nums); currSum := 0
    for left < n {
        if right >= n {
            for currSum >= target && left < n {
                minimum = min(minimum, n-left)
                currSum -= nums[left]
                left++
            }
            break
        }

        currSum += nums[right]
        if currSum < target {
            right++
            continue
        }

        if currSum >= target {
            for currSum >= target && left < n {
                minimum = min(minimum, right-left+1)
                currSum -= nums[left]
                left++
            }
            right++
        }
    }

    if minimum == math.MaxInt64 {
        return 0
    }

    return minimum
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

func twoSum(numbers []int, target int) []int {
    ans := make([]int, 2)
    n := len(numbers)

    isFound := false
    for i, number := range numbers {
        pair := target - number
        left := i + 1
        right := n - 1
        for left <= right {
            mid := (left + right) / 2
            if numbers[mid] == pair {
                ans[0] = i+1
                ans[1] = mid+1
                isFound = true
                break
            }
            if (numbers[mid] < pair) {
                left = mid + 1
            }
            if (numbers[mid] > pair) {
                right = mid - 1
            }
        }
        if isFound {
            break
        }
    }

    return ans
}
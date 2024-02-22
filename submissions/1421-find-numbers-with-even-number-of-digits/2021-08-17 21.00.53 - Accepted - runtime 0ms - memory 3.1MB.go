func countDigit(num int) int {
    if num == 0 {
        return 1
    }
    
    cnt := 0
    for num > 0 {
        cnt += 1
        num /= 10
    }
    return cnt
}

func findNumbers(nums []int) int {
    cnt := 0
    for _, num := range nums {
        if countDigit(num) % 2 == 0 {
            cnt += 1
        }
    }
    return cnt
}
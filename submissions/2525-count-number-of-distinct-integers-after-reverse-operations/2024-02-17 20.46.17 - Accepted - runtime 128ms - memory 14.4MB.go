func countDistinctIntegers(nums []int) int {
    mep := make(map[int]bool)
    for _, num := range nums {
        mep[num] = true
    }
    for _, num := range nums {
        reversed := reverse(num)
        mep[reversed] = true
    }

    return len(mep)
}

func reverse(num int) int {
    r := 0
    for num > 0 {
        r = 10 * r + num % 10
        num /= 10
    }

    return r
}
func singleNumber(nums []int) int {
    xorVal := 0
    for _, num := range nums {
        xorVal ^= num
    }
    return xorVal
}
func isUgly(n int) bool {
    if n <= 0 {
        return false
    }
    
    facNums := [3]int{2, 3, 5}
    for _, num := range facNums {
        for n % num == 0{
            n /= num
        }
    }
    return n == 1
}
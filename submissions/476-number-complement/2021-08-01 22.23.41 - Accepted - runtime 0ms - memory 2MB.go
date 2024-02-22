func findComplement(num int) int {
    reverse := 0
    multiplier := 1
    
    for num > 0 {
        if num % 2 == 0 {
            reverse += multiplier
        }
        multiplier *= 2
        num /= 2
    }
    
    return reverse
}
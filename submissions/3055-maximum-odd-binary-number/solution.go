func maximumOddBinaryNumber(s string) string {
    ones := 0
    for _, ch := range s {
        if ch == '1' {
            ones++
        }
    }

    ans := ""
    for i := 0; i < ones-1; i++ {
        ans = ans + "1"
    }
    for i := 0; i < len(s)-ones; i++ {
        ans = ans + "0"
    }
    ans = ans + "1"

    return ans
}

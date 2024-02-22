func balancedStringSplit(s string) int {
    ans := 0; countL := 0; countR := 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'L' {
            countL += 1
        } else {
            countR += 1
        }
        if countL == countR {
            ans += 1
            countL = 0
            countR = 0
        }
    }
    return ans
}
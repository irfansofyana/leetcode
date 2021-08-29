func minDist(a byte, b byte) int {
    diff := int(a) - int(b)
    if diff < 0 {
        diff *= -1
    }
    if 26-diff < diff {
        diff = 26-diff
    }
    return diff
}

func minTimeToType(word string) int {
    now := byte('a')
    
    ans := 0
    for i := 0; i < len(word); i++ {
        ans += minDist(now, byte(word[i])) + 1
        now = byte(word[i])
    }
    
    return ans
}

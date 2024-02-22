func areOccurrencesEqual(s string) bool {
    freq := map[int]int{}; n := len(s)
    for i := 0; i < n ; i++ {
        ch := int(s[i] - 'a')
        freq[ch] += 1
    }
    ref := -1
    for _, v := range freq {
        if ref == -1  {
            ref = v
        } else if ref != v {
            return false
        }
    }
    return true
}
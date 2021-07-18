func minSteps(s string, t string) int {
    var strLen int = len(s)
    var freq1 = make([]int, 26)
    var freq2 = make([]int, 26)
    
    for i := 0; i < strLen; i++ {
        freq1[s[i] - 'a'] += 1
        freq2[t[i] - 'a'] += 1
    }
    
    for i := 0; i < 26; i++ {
        if freq1[i] <= freq2[i] {
            strLen -= freq1[i]
        } else {
            strLen -= freq2[i]
        }
    }
    
    return strLen
}

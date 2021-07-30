func isAnagram(s string, t string) bool {
    freqs := make([]int, 26); freqt := make([]int, 26)
    for i := 0; i < len(s); i++ {
        freqs[s[i]-'a'] += 1
    }
    for i := 0; i < len(t); i++ {
        freqt[t[i]-'a'] += 1
    }
    
    isanagram := true
    for i := 0; i < 26; i++ {
        if freqs[i] != freqt[i] {
            isanagram = false
            break
        }
    } 
    return isanagram
}

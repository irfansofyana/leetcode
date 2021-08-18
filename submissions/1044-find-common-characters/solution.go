func commonChars(words []string) []string {
    n := len(words)
    freq := make([][26]int, n)
    
    for i := 0; i < n; i++ {
        for j := 0; j < len(words[i]); j++ {
            ch := int(words[i][j] - 'a')
            freq[i][ch] += 1
        }
    }
    
    ans := make([]string, 0)
    for i := 0; i < 26; i++ {
        mini := 10000
        for j := 0; j < n; j++ {
            if freq[j][i] < mini {
                mini = freq[j][i]
            }
        }
        for j := 0; j < mini; j++ {
            ans = append(ans, string(int('a') + i))
        }
    }
    
    return ans
}

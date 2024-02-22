func findTheDifference(s string, t string) byte {
    frek := make([]int, 26)
    ns := len(s); nt := len(t)
    var ans int = 0
    
    for i := 0; i < nt; i++ {
        frek[t[i]-'a'] += 1
    }
    for i := 0; i < ns; i++ {
        frek[s[i]-'a'] -= 1
    }
    for i := 0; i < 26; i++ {
        if frek[i] == 1 {
            ans = i
            break
        }
    } 
    return byte(ans + 97)
}
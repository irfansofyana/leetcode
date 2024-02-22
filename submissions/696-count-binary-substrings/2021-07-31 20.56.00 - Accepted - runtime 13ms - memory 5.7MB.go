func countBinarySubstrings(s string) int {
    n := len(s); ans := 0; lst := -1;
    
    i := 0
    for i < n {
        j := i + 1
        curr := s[i]
        frek := 1
        for j < n && s[j] == curr {
            frek += 1
            j += 1
        }
        if lst == -1 {
            lst = frek
        } else {
            if lst < frek {
                ans += lst
            } else {
                ans += frek
            }
            lst = frek
        }
        i = j
    }
    
    return ans
}
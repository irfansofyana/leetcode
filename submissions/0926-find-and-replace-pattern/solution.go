func isValid(a string, b string) bool {
    if len(a) != len(b) {
        return false
    }
    usedA := map[int]int{}
    usedB := map[int]int{}
    for i := 0; i < len(a); i++ {
        chb := int(b[i])
        cha := int(a[i])
        if usedA[cha] > 0 && usedA[cha] != chb {
            return false
        }
        if usedB[chb] > 0 && usedB[chb] != cha {
            return false
        }
        usedA[cha] = chb
        usedB[chb] = cha
    }
    return true
}

func findAndReplacePattern(words []string, pattern string) []string {
    ans := make([]string, 0)
    for _, word := range words {
        if isValid(word, pattern) {
            ans = append(ans, word)
        }
    }
    return ans
}

func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word) {
            return word
        }
    }

    return ""
}

func isPalindrome(st string) bool {
    l := 0; r := len(st)-1

    for l <= r {
        if st[l] != st[r] {
            return false
        }
        l++; r--
    }

    return true
}

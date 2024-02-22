func isPalindrome(s string, i, j int) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++; j--
    }

    return true
}

func validPalindrome(s string) bool {
    left := 0; right := len(s) - 1

    for left < right {
        if s[left] != s[right] {
            return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
        } else {
            left++
            right--
        }
    }

    return true
}
var ans [][]string

func partition(s string) [][]string {
    ans = make([][]string, 0)

    generate(s, 1, 1)    

    return ans
}

func generate(s string, pos int, bitmask int) {
    if pos == len(s) {
        curr := ""; cumm := make([]string, 0)
        i := 0
        for i < pos {
            if (bitmask&(1<<i) > 0) {
                curr = string(s[i])
                i++
                for i < pos && bitmask&(1<<i) == 0 {
                    curr += string(s[i])
                    i++
                }
                if (!isPalindrome(curr)) {
                    return
                }
                cumm = append(cumm, curr)
            }
        }

        ans = append(ans, cumm)
        return
    }

    generate(s, pos + 1, bitmask)

    generate(s, pos + 1, bitmask|(1<<pos))
}

func isPalindrome(s string) bool {
    l := 0; r := len(s)-1
    for l <= r {
        if s[l] != s[r] {
            return false
        }

        l++; r--
    }

    return true
}


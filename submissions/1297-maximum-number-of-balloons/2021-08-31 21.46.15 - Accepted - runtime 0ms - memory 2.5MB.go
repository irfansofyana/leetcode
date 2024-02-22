func findIdx(ch rune) int {
    st := "balon"
    ans := -1
    for i, c := range st {
        if c == ch {
            ans = i
            break
        }
    }
    return ans
}

func maxNumberOfBalloons(text string) int {
    freq := make([]int, 6)
    for _, t := range text {
        idx := findIdx(t)
        if idx != -1{
            freq[idx]++   
        }
    }
    ans := 1000000000
    for i := 0; i < 5; i++ {
        var cnt int = freq[i]
        if i == 2 || i == 3 {
            cnt /= 2
        }
        if cnt < ans {
            ans = cnt
        }
    }
    return ans
}
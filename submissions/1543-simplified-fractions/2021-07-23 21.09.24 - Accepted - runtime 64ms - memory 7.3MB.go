func gcd(a int, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func simplifiedFractions(n int) []string {
    var ans = make([]string, 0)
    
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            if gcd(i, j) == 1 {
                ans = append(ans, fmt.Sprintf("%d/%d", j, i))
            }
        }
    }
    
    return ans
}
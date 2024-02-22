import "math"

func judgeSquareSum(c int) bool {
    if c == 0 {
        return true
    }
    
    yes := false
    for b := 1; b <= c / b; b++ {
        a := c - b * b
        sqrtA := int(math.Sqrt(float64(a)))
        if sqrtA * sqrtA == a {
            yes = true
            break
        }
    }
    return yes
}
import "sort"

func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    score := 0
    l := 0; r := len(tokens)-1
    for l <= r {
        if l == r {
            if power >= tokens[l] {
                power -= tokens[l]
                score++
            }
            break
        }

        if power >= tokens[l] {
            score++
            power -= tokens[l]
            l++
        } else if power < tokens[l] && score > 0 {
            power += tokens[r]
            score--
            r--
        } else {
            l++
        }   
    }
    return score
}

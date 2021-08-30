import "sort"

func maxCoins(piles []int) int {
    n := len(piles)
    sort.Ints(piles)
    ans := 0
    for i := n / 3; i < n; i+=2 {
        ans += piles[i]
    }
    return ans
}


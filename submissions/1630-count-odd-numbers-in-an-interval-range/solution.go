func countOdds(low int, high int) int {
    cnt := func (n int) int {
        if n % 2 == 1 {
            return n / 2 + 1
        }
        return n / 2
    }
    ans := cnt(high)
    if low > 1 {
        ans -= cnt(low-1)
    }
    return ans
}

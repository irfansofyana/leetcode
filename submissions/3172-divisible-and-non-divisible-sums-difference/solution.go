func differenceOfSums(n int, m int) int {
    total := n * (n + 1) / 2
    divisible := m * ((n / m) * (n / m + 1)) / 2
    return total - 2 * divisible
}

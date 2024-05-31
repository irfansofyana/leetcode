func maxProfit(prices []int, fee int) int {
    cash := 0; hold := -2000000000
    for i := 0; i < len(prices); i++ {
        cash = max(cash, hold + prices[i] - fee)
        hold = max(hold, cash - prices[i])
    }

    return cash
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func maxProfit(prices []int) int {
    profitMax := 0
    n := len(prices)
    maxPrice := prices[n-1]
    
    for i := n-2; i >= 0; i-- {
        if maxPrice - prices[i] > profitMax {
            profitMax = maxPrice - prices[i]
        }
        if prices[i] > maxPrice {
            maxPrice = prices[i]
        }
    }
    
    return profitMax
}
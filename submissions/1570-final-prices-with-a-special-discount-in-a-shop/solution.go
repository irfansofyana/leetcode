func finalPrices(prices []int) []int {
    n := len(prices)
    final := make([]int, n)
    
    for i := 0; i < n; i++ {
        idx := -1
        for j := i+1; j < n; j++ {
            if prices[j] <= prices[i] {
                idx = j
                break
            }
        }
        discount := 0
        if idx != -1{
            discount = prices[idx]
        }
        final[i] = prices[i] - discount
    }
    
    return final
}

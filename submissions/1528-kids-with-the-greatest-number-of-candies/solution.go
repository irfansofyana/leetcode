func kidsWithCandies(candies []int, extraCandies int) []bool {
    ans := make([]bool, 0); n := len(candies)
    
    for i := 0; i < n; i++ {
        isValid := true
        for j := 0; j < n && isValid; j++ {
            if candies[j] > candies[i] + extraCandies {
                isValid = false
            }
        }
        ans = append(ans, isValid)
    }
    
    return ans
}

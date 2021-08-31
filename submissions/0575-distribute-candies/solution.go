func distributeCandies(candyType []int) int {
    n := len(candyType); limit := n / 2
    freqType := map[int]int{}
    
    for _, t := range candyType {
        freqType[t]++
    }
    
    distinctType := len(freqType)
    if distinctType < limit {
        return distinctType
    }
    return limit
}

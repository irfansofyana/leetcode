func groupThePeople(groupSizes []int) [][]int {
    var freq = map[int][]int{}
    var ans = make([][]int, 0)
    
    for i, groupSize := range groupSizes {
        if len(freq[groupSize]) == 0 {
            freq[groupSize] = make([]int, 0)
        }
        freq[groupSize] = append(freq[groupSize], i)
    }
    
    for k, peoples := range freq {
        var numGroup = len(peoples) / k
        for i := 0; i < numGroup; i++ {
            ans = append(ans, peoples[i*k:i*k+k])
        }
    }
    
    return ans
}
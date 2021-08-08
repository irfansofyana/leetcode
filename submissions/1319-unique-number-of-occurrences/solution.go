func uniqueOccurrences(arr []int) bool {
    freq := make([]map[int]int, 2); n := len(arr)
    for i := 0; i < 2; i++ {
        freq[i] = map[int]int{}
    }
    
    for i := 0; i < n; i++ {
        freq[0][arr[i]] += 1
    }
    
    isUnique := true
    for _, v := range freq[0] {
        if freq[1][v] > 0 {
            isUnique = false
            break
        }
        freq[1][v] += 1
    }
    
    return isUnique
}

func numSplits(s string) int {
    n := len(s); forward := make([]int, 0); backward := make([]int, 0)
    
    freq := map[int]int{}; unique := 0
    for i := 0; i < n; i++ {
        idx := int(s[i] - 'a')
        freq[idx] += 1
        if freq[idx] == 1 {
            unique += 1
        }
        forward = append(forward, unique)
    }
    
    unique = 0; freq = map[int]int{}
    for i := n-1; i >= 0; i-- {
        idx := int(s[i] - 'a')
        freq[idx] += 1
        if freq[idx] == 1 {
            unique += 1
        }
        backward = append(backward, unique)
    }
    
    cnt := 0
    for i := 0; i < n-1; i++ {
        left := forward[i]
        right := backward[n-2-i]
        
        if left == right {
            cnt += 1
        }
    }
    
    return cnt
}

func findRemainder(n int, k int) int {
    if n >= 0 {
        return n % k
    }
    tmp := (-1 * n) % k
    return (k - tmp)%k
}

func canArrange(arr []int, k int) bool {
    freq := map[int]int{}
    n := len(arr)
    can := true
    
    for i := 0; i < n; i++ {
        freq[findRemainder(arr[i], k)] += 1
    }
    for i := 0; i < k; i++ {
        if findRemainder(i, k) == findRemainder(k-i, k) && freq[findRemainder(i, k)] % 2 == 1 {
            can = false
            break
        }
        if freq[findRemainder(i, k)] != freq[findRemainder(k-i, k)] {
            can = false
            break
        }
    }
    
    return can
}

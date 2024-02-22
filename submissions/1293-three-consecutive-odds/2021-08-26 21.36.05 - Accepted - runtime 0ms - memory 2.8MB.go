func threeConsecutiveOdds(arr []int) bool {
    var i int
    n := len(arr)
    for i < n {
        for i < n && arr[i] % 2 == 0 {
            i++
        }
        curr := 1
        j := i + 1
        for j < n && arr[j] % 2 == 1 {
            curr++
            if curr == 3 {
                return true
            }
            j++
        }
        i = j
    }
    return false
}
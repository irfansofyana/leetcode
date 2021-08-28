func nextGreatestLetter(letters []byte, target byte) byte {
    n := len(letters)
    ans := n
    l := 0
    r := n - 1
    
    for l <= r {
        mid := (l + r) / 2
        if letters[mid] > target {
            if mid < ans {
                ans = mid
            }
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    
    if ans == n {
        return letters[0]
    }
    return letters[ans]
}

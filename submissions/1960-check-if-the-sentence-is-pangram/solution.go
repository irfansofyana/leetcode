func checkIfPangram(sentence string) bool {
    freq := make([]int, 26); n := len(sentence)
    
    for i := 0; i < n; i++ {
        idx := int(sentence[i] -'a')
        freq[idx] += 1
    }
    
    isPangram := true
    for i := 0; i < 26; i++ {
        if freq[i] == 0{
            isPangram = false
            break
        }
    }
    
    return isPangram
}

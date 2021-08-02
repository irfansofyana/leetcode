func detectCapitalUse(word string) bool {
    var isFirstWordCapital bool = false
    if word[0] >= 'A' && word[0] <= 'Z' {
        isFirstWordCapital = true
    }
    
    isTrue := true
    for i := 1; i < len(word); i++ {
        if !isFirstWordCapital && (word[i] >= 'A' && word[i] <= 'Z' && i < len(word)) {
            isTrue = false
            break
        }
        if isFirstWordCapital && (word[1] >= 'A' && word[1] <= 'Z') && (i > 1 && word[i] >= 'a' && word[i] <= 'z') {
            isTrue = false
            break
        }
        if isFirstWordCapital && (word[1] >= 'a' && word[1] <= 'z') && (i > 1 && word[i] >= 'A' && word[i] <= 'Z') {
            isTrue = false
            break
        }
    }
    
    return isTrue
}

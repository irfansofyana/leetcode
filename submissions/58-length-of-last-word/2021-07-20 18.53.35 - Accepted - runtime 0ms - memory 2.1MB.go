func lengthOfLastWord(s string) int {
    var lastWordLen int = 0
    var idx int = 0
    var curr int = 0
    
    for idx < len(s) {
        for idx < len(s) && s[idx] == ' ' {
            if curr > 0 {
                lastWordLen = curr
                curr = 0
            }
            idx += 1
        }
        if idx >= len(s) {
            break;
        }
        curr += 1
        idx += 1
    }
    if curr > 0 {
        lastWordLen = curr
    }
    
    return lastWordLen
}
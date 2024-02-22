func isMakeSubtraction(a string, b string) bool {
    if a == "I" && (b == "V" || b == "X") {
        return true
    }
    if a == "X" && (b == "L" || b == "C") {
        return true
    }
    if a == "C" && (b == "D" || b == "M") {
        return true
    }
    return false
}

func romanToInt(s string) int {
    var convertedInt int = 0
    var pos int = 0
    var charMap = map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    
    for pos < len(s) {
        if s[pos] == 'I' || s[pos] == 'X' || s[pos] == 'C' {
            if pos + 1 < len(s) &&  isMakeSubtraction(string(s[pos]), string(s[pos+1])) {
                convertedInt += charMap[string(s[pos+1])] - charMap[string(s[pos])]
                pos += 2
                continue
            }
        }
        convertedInt += charMap[string(s[pos])]
        pos += 1
    }
    
    return convertedInt
}
func letterCombinations(digits string) []string {
    mapDigits := buildMapping(); combinations := make([]string, 0)
    
    generateCombinations(0, "", digits, mapDigits, &combinations)

    return combinations
}

func generateCombinations(pos int, currString string, digits string, mapDigits map[string]string, result *[]string) {
    if pos == len(digits) {
        if currString != "" {
            *result = append(*result, currString)
        }
        return
    }

    str := mapDigits[string(digits[pos])]
    for i := 0; i < len(str); i++ {
        generateCombinations(pos+1, currString+string(str[i]), digits, mapDigits, result)
    }
}

func buildMapping() map[string]string {
    mapDigits := make(map[string]string)
    mapDigits["2"] = "abc"
    mapDigits["3"] = "def"
    mapDigits["4"] = "ghi"
    mapDigits["5"] = "jkl"
    mapDigits["6"] = "mno"
    mapDigits["7"] = "pqrs"
    mapDigits["8"] = "tuv"
    mapDigits["9"] = "wxyz"
    
    return mapDigits
}

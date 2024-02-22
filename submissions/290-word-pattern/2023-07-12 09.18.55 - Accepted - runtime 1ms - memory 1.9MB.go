import "strings"

func wordPattern(pattern string, s string) bool {
    patternMapping := make(map[byte]string)
    wordMapping := make(map[string]byte)
    splitted := strings.Split(s, " ")

    if len(pattern) != len(splitted) {
        return false
    }
    
    for i := 0; i < len(splitted); i++ {
        if _, ok := patternMapping[pattern[i]]; !ok {
            if _, okk := wordMapping[splitted[i]]; !okk {
                patternMapping[pattern[i]] = splitted[i]
                wordMapping[splitted[i]] = pattern[i]
                continue
            }

            return false
        }

        if patternMapping[pattern[i]] != splitted[i] {
            return false
        }

        if wordMapping[splitted[i]] != pattern[i] {
            return false
        }
    }

    return true
}
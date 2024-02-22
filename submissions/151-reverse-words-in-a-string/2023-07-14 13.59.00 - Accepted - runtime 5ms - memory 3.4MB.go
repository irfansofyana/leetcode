import (
    "strings"
    "regexp"
)

func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    whitespaces := regexp.MustCompile(`\s+`)
    s = whitespaces.ReplaceAllString(s, " ")    
    
    splitted := strings.Split(s, " ")
    for i, j := 0, len(splitted)-1; i < j; i, j = i + 1, j - 1 {
        splitted[i], splitted[j] = splitted[j], splitted[i]
    }
    
    return strings.Join(splitted, " ")    
}
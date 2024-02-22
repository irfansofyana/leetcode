import "strings"

func checkRecord(s string) bool {
    return strings.Count(s,"A") < 2 && strings.Count(s, "LLL") == 0
}
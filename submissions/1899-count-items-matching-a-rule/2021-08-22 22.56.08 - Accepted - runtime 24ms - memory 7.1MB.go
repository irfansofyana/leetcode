func ruleToIdx(rule string) int {
    if rule == "type" {
        return 0
    }
    if rule == "color" {
        return 1
    }
    return 2
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    ans := 0; idx := ruleToIdx(ruleKey)
    for _, item := range items {
        if item[idx] == ruleValue {
            ans += 1
        }
    }
    return ans
}
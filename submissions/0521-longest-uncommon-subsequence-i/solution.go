func findLUSlength(a string, b string) int {
    la := len(a)
    lb := len(b)

    if la != lb {
        return max(la, lb)
    }
    if a == b {
        return -1
    }

    return la
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

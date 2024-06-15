func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    parent := make([]int, 26)
    for i := 0; i < 26; i++ {
        parent[i] = i
    }    

    n := len(s1)
    for i := 0; i < n; i++ {
        union(&parent, int(s1[i])-'a', int(s2[i])-'a')
    }

    ans := ""
    for i := 0; i < len(baseStr); i++ {
        parentChar := findParent(&parent, int(baseStr[i])-'a')
        ans += string(int('a') + parentChar)
    }

    return ans
} 

func union(parent *[]int, a int, b int) {
    pa := findParent(parent, a)
    pb := findParent(parent, b)
    if pa < pb {
        (*parent)[pb] = pa
    } else {
        (*parent)[pa] = pb
    }
}

func findParent(parent *[]int, a int) int {
    if len(*parent) <= a {
        return -1
    }

    if a == (*parent)[a] {
        return a
    }

    (*parent)[a] = findParent(parent, (*parent)[a])
    return (*parent)[a]
}

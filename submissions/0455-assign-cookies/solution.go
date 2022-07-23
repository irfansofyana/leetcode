func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    
    content := 0
    gIndex := 0; sIndex := 0
    for sIndex < len(s) && gIndex < len(g) {
        if s[sIndex] >= g[gIndex] {
            content++
            sIndex++
            gIndex++
            continue
        }
        
        sIndex++
    }
    
    return content
}

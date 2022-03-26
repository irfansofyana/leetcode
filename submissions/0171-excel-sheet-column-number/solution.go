func titleToNumber(columnTitle string) int {
    num := 0
    pow := 1
    
    for i := len(columnTitle)-1; i >= 0; i-- {
        num += pow * (int(columnTitle[i])-64)
        pow *= 26
    }
    
    return num
}

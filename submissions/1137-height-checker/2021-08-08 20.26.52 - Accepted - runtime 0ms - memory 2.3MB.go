import "sort"

func heightChecker(heights []int) int {
    tmp := make([]int, 0)
    tmp = append(tmp, heights...)
    
    sort.Ints(tmp)
    
    cnt := 0
    for i, height := range heights {
        if height != tmp[i] {
            cnt += 1
        } 
    }
    
    return cnt
}
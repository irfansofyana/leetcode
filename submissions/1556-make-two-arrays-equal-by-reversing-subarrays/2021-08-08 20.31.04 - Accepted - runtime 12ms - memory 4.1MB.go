import "sort"

func canBeEqual(target []int, arr []int) bool {
    sort.Ints(target); sort.Ints(arr)
    
    isCanEqual := true
    for i, num := range arr {
        if num != target[i] {
            isCanEqual = false
            break
        }
    }
    
    return isCanEqual
}
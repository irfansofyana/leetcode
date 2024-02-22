import "sort"

func minimumAbsDifference(arr []int) [][]int {
    f := map[int][][]int{}
    
    sort.Ints(arr); n := len(arr); mini := 2000000000
    for i := 0; i < n-1; i++ {
        diff := arr[i+1] - arr[i]
        if _, ok := f[diff]; ok {
            f[diff] = append(f[diff], []int{arr[i], arr[i+1]})
        } else {
            f[diff] = make([][]int, 0)
            f[diff] = append(f[diff], []int{arr[i], arr[i+1]})
        }
        if diff < mini {
            mini = diff
        }
    }
    
    return f[mini]
}
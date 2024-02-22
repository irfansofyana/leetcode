import "sort"

func hIndex(citations []int) int {
    sort.Ints(citations)
    n := len(citations)

    ans := 0
    for i := n-1; i >= 0; i-- {
        num := n-i
        if citations[i] >= num {
            ans = num
            continue
        } else {
            break
        }
    }

    return ans    
}
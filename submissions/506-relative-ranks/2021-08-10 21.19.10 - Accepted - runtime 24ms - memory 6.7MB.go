import (
    "fmt"
    "sort"
)

type dat struct {
    order int
    score int
}

func findRelativeRanks(score []int) []string {
    n := len(score)
    
    arr := make([]dat, 0)
    for i := 0; i < n; i++ {
        arr = append(arr, dat{order: i, score: score[i]})
    }
    
    sort.SliceStable(arr, func (i, j int) bool {
        return arr[i].score > arr[j].score
    })
    
    ans := make([]string, n)
    for i := 0; i < n; i++ {
        var award string = fmt.Sprintf("%d", (i+1))
        if i == 0 {
            award = "Gold Medal"
        } else if i == 1 {
            award = "Silver Medal"
        } else if i == 2 {
            award = "Bronze Medal"
        }
        ans[arr[i].order] = award
    }
    
    return ans
}
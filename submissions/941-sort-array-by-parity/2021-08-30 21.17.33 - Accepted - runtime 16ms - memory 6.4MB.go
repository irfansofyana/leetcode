import "sort"

type Num struct {
    Val int    
}

func sortArrayByParity(nums []int) []int {
    arr := make([]Num, 0)
    for _, num := range nums {
        arr = append(arr, Num{Val: num})
    }
    sort.SliceStable(arr, func(i int, j int) bool {
        if arr[i].Val % 2 == arr[j].Val % 2 {
            return arr[i].Val < arr[j].Val
        }
        return (arr[i].Val % 2) < (arr[j].Val % 2)
    })
    ans := make([]int, 0)
    for i := 0; i < len(arr); i++ {
        ans = append(ans, arr[i].Val)
    }
    return ans
}
import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
    freq := map[int]int{}
    for _, num := range arr1 {
        freq[num]++
    }
    
    ans := make([]int, 0)
    for _, num := range arr2 {
        for i := 0; i < freq[num]; i++ {
            ans = append(ans, num)
        }
        freq[num] = 0
    }
    
    t := make([]int, 0)
    for k, v := range freq {
        for j := 0; j < v; j++ {
            t = append(t, k)
        }
    }
    
    sort.Ints(t)
    ans = append(ans, t...)
    
    return ans
}
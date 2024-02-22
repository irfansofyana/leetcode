func intersect(nums1 []int, nums2 []int) []int {
    freq := make([]map[int]int, 2)
    arr := [2][]int{nums1, nums2}
    
    for i := 0; i < 2; i++ {
        freq[i] = map[int]int{}
        for _, num := range arr[i] {
            freq[i][num]++
        }
    }
    
    ans := make([]int, 0)
    for k, v := range freq[0] {
        if freq[1][k] == 0 {
            continue
        }
        mini := v
        if freq[1][k] < mini {
            mini = freq[1][k]
        }
        for j := 0; j < mini; j++ {
            ans = append(ans, k)
        }
    }
    
    return ans
}
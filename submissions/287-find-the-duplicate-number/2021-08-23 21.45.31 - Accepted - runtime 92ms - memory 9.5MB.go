func findDuplicate(nums []int) int {
    n := len(nums)-1; freq := make([]int, n+1)
    var ans int
    for _, num := range nums {
        freq[num]++
        if freq[num] > 1 {
            ans = num
        }
    }
    return ans
}
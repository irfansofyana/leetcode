func majorityElement(nums []int) int {
    var freq = map[int]int{}
    var ans int
    
    for _, num := range nums {
        freq[num] += 1
        if freq[num] > len(nums) / 2{
            ans = num
            break
        }
    }
    
    return ans
}

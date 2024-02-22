func findLucky(arr []int) int {
    freq := map[int]int{}
    for _, num := range arr {
        freq[num]++
    }
    ans := -1
    for num, cnt := range freq {
        if num == cnt  && num > ans {
            ans = num
        }
    }
    return ans
}
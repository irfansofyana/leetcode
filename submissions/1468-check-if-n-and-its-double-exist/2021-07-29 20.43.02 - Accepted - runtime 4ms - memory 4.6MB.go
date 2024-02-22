func checkIfExist(arr []int) bool {
    isExist := map[int]int{}
    ans := false
    
    for _, num := range arr {
        isExist[num] += 1
    }
    for _, num := range arr {
        if (num != 2 * num && isExist[2*num] > 0) || (num == 2 * num && isExist[2*num] >= 2 ) {
            ans = true
            break
        }
    }
    return ans
}
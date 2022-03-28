func sumArr(arr []int) int {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    return sum
}

func chalkReplacer(chalk []int, k int) int {
    sum := sumArr(chalk)
    rem := k % sum
    
    var ans int
    for i, num := range chalk {
        if rem - num < 0 {
            ans = i
            break
        } 
        rem -= num
    }
    
    return ans
}

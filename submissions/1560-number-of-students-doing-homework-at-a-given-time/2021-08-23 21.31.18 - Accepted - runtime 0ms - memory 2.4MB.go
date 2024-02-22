func busyStudent(startTime []int, endTime []int, queryTime int) int {
    n := len(startTime)
    
    ans := 0
    for i := 0; i < n; i++ {
        if startTime[i] <= queryTime && queryTime <= endTime[i] {
            ans += 1
        }    
    }
    
    return ans
}
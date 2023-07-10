func largestAltitude(gain []int) int {
    largest := 0
    currSum := 0
    
    for _, height := range gain {
        currSum += height
        if currSum > largest {
            largest = currSum
        }
    }

    return largest
}

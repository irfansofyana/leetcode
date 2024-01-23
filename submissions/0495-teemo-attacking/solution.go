import "sort"

func findPoisonedDuration(timeSeries []int, duration int) int {
    sort.Ints(timeSeries)

    total := duration; 
    currLeft := timeSeries[0]; 
    currRight := timeSeries[0] + duration - 1

    for i := 1; i < len(timeSeries); i++ {
        if timeSeries[i] > currRight {
            total += duration
        } else {
            total -= (currRight - timeSeries[i] + 1)
            total += duration
        }
        currLeft = timeSeries[i]
        currRight = currLeft + duration - 1
    }

    return total    
}

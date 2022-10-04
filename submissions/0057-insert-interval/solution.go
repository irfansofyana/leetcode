import "sort"


func insert(intervals [][]int, newInterval []int) [][]int {
    mergedIntervals := make([][]int, 0)
    
    intervals = append(intervals, newInterval)
    
    sort.SliceStable(intervals, func (i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        
        return intervals[i][0] < intervals[j][0]
    })
    
    if len(intervals) == 0 {
        return mergedIntervals
    }
    
    currLeft := intervals[0][0]
    currRight := intervals[0][1]
    
    i := 1
    for i < len(intervals) {
        interval := intervals[i]
        
        if interval[0] > currRight {
            mergedIntervals = append(mergedIntervals, []int{currLeft, currRight})
            currLeft = interval[0]
            currRight = interval[1] 
            i++
            continue
        }
        
        if interval[0] <= currRight {
            if interval[1] > currRight {
                currRight = interval[1]
            }
            i++
        }
    }
    
    mergedIntervals = append(mergedIntervals, []int{currLeft, currRight})
    
    return mergedIntervals
}

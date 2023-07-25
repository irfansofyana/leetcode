import "sort"

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }

        return intervals[i][0] < intervals[j][0]
    })

    merged := make([][]int, 0)
    currLeft := intervals[0][0]
    currRight := intervals[0][1]

    i := 1
    for i < len(intervals) {
        lPos := intervals[i][0]
        rPos := intervals[i][1]
        if lPos > currRight {
            merged = append(merged, []int{currLeft, currRight})
            currLeft = lPos
            currRight = rPos
            i++
            continue
        }

        if lPos <= currRight && rPos <= currRight {
            i++
            continue
        }

        if lPos <= currRight && rPos > currRight {
            currRight = rPos
            i++
            continue
        }
    }
    merged = append(merged, []int{currLeft, currRight})

    return merged
}

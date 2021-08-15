import (
    "strconv"
    "sort"
)

func convertToNumber(times string) int {
    hour := times[0:2]; minute := times[3:]
    hhour, _ := strconv.Atoi(hour); 
    mminute, _ := strconv.Atoi(minute)
    return 60 * hhour + mminute
}

func diffTime(num1, num2 int) int {
    var diff int
    if num1 > num2 {
        diff = num1 - num2
    } else {
        diff = num2 - num1
    }
    
    if 24*60 - diff < diff {
        diff = 24*60 - diff
    }
    
    return diff
}

func findMinDifference(timePoints []string) int {
    arr := make([]int, 0); n := len(timePoints)
    for i := 0; i < n; i++ {
        arr = append(arr, convertToNumber(timePoints[i]))
    }
    sort.Ints(arr)
    ans := 10000
    for i := 0; i < n; i++ {
        var diff int
        if i == n-1 {
            diff = diffTime(arr[n-1], arr[0])
        } else {
            diff = diffTime(arr[i+1], arr[i])    
        }
        if diff < ans {
            ans = diff
        }
    }
    return ans
}

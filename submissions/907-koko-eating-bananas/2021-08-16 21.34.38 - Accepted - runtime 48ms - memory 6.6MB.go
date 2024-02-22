func can(piles []int, limitTime int, givenTime int) bool {
    totalTime := 0
    for _, pile := range piles {
        totalTime += (pile + givenTime - 1) / givenTime
    }
    return totalTime <= limitTime
}

func minEatingSpeed(piles []int, h int) int {
   lo := 1; hi := 2000000000
    
    ans := hi
    for lo <= hi {
        mid := (lo + hi) / 2
        if can(piles, h, mid) {
            ans = mid
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    
    return ans
}
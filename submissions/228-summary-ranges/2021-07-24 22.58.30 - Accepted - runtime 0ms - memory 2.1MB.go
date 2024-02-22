import "fmt"

func summaryRanges(nums []int) []string {
    ans := make([]string, 0)
    idx := 0
    n := len(nums)
    
    for idx < n {
        nextIdx := idx + 1
        for nextIdx < n && nums[nextIdx] - nums[nextIdx-1] == 1 {
            nextIdx += 1
        }
        if nextIdx == idx + 1 {
            ans = append(ans, fmt.Sprintf("%d", nums[idx]))
        } else {
            ans = append(ans, fmt.Sprintf("%d->%d", nums[idx], nums[nextIdx-1]))
        }
        idx = nextIdx
    }
    
    return ans
}
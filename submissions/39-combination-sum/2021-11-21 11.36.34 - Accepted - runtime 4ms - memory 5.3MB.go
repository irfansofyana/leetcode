import "sort"

func combinationSum(candidates []int, target int) [][]int {
    dp := make([][][]int, target+1)
    
    sort.Ints(candidates)
    for i := 1; i <= target; i++ {
        for j := 0; j < len(candidates); j++ {
            if i >= candidates[j] {
                rem := i - candidates[j]
                for k := 0; k < len(dp[rem]); k++ {
                    if candidates[j] >= dp[rem][k][0] {
                        tmp := append([]int{candidates[j]}, dp[rem][k]...)
                        dp[i] = append(dp[i], tmp)
                    }
                }
                if rem == 0 {
                    tmp := []int{candidates[j]}
                    dp[i] = append(dp[i], tmp)
                }
            }
        }    
    }
    
    
    return dp[target]
}
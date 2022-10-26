/*
    prefix[i]-prefix[j-1] == 0 mod k where i!=j
    5 2 0 4 1
    5 1 1 5 0
*/


func checkSubarraySum(nums []int, k int) bool {
    prefixSum := make([]int, len(nums))
    
    for i, num := range nums {
        num %= k
        
        if i == 0 {
            prefixSum[i] = num
            continue
        }
        
        prefixSum[i] = (num + prefixSum[i-1]) % k
    }
    
    mepMod := make(map[int][]int)
    
    for i, p := range prefixSum {
        mepMod[p] = append(mepMod[p], i)
    }
    
    for mod := range mepMod {
        arr := mepMod[mod]
        lenArr := len(arr)
        
        if mod == 0 && arr[lenArr-1] != 0 {
            return true
        }
        
        if lenArr == 1 {
            continue
        }
        
        if arr[lenArr-1] - arr[0] == 1 {
            continue
        }
        
        
        return true
    }
    
    return false
}

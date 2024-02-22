type NumArray struct {
    prefixSum []int
}


func Constructor(nums []int) NumArray {
    prefixSum := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            prefixSum[i] = nums[i]
            continue
        }
        
        prefixSum[i] = prefixSum[i-1] + nums[i]
    }
    
    return NumArray{prefixSum}
}


func (this *NumArray) SumRange(left int, right int) int {
    var sub int = 0
    if left != 0 {
        sub = this.prefixSum[left-1] 
    }
    return this.prefixSum[right] - sub
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
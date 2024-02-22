/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    idxMaks := 0; maks := nums[0]; n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] > maks {
            idxMaks = i
            maks = nums[i]
        }
    }
    
    if n == 1 {
        return &TreeNode{maks, nil, nil}        
    }
    
    root := &TreeNode{
        maks,
        constructMaximumBinaryTree(nums[0:idxMaks]),
        constructMaximumBinaryTree(nums[idxMaks+1:n]),
    }
    return root
}
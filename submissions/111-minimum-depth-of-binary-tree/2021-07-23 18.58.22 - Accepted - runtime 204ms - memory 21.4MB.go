/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    minLeft := minDepth(root.Left) + 1
    minRight := minDepth(root.Right) + 1
    
    if root.Left == nil {
        return minRight
    }
    if root.Right == nil {
        return minLeft
    }
    
    if minLeft > minRight {
        minLeft, minRight = minRight, minLeft
    }
    
    return minLeft
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
    if a > b {
        return a
    }    
    return b
}

func countHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return max(countHeight(root.Left), countHeight(root.Right)) + 1
}

func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    leftH := countHeight(root.Left)
    rightH := countHeight(root.Right)
    
    if abs(leftH, rightH) <= 1 {
        return isBalanced(root.Left) && isBalanced(root.Right)
    }
    
    return false
}
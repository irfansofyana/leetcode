/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftHeight := height(root.Left)
    rightHeight := height(root.Right)
    
    leftDiameter := diameterOfBinaryTree(root.Left)
    rightDiameter := diameterOfBinaryTree(root.Right)
    
    return max(max(leftDiameter, rightDiameter), leftHeight+rightHeight)
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftHeight := height(root.Left)
    rightHeight := height(root.Right)
    
    return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}

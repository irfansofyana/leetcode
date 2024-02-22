/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    maksLeft := maxDepth(root.Left) 
    maksRight := maxDepth(root.Right)
    if maksLeft > maksRight {
        return maksLeft + 1
    }
    return maksRight + 1
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
    leftNode:= root.Left
    rightNode:= root.Right
    
    return root.Val == leftNode.Val + rightNode.Val
}

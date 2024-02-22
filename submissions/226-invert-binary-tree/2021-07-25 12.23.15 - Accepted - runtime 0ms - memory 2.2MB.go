/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var ret *TreeNode = &TreeNode{root.Val, nil, nil}
    leftTree := invertTree(root.Left)
    rightTree := invertTree(root.Right)
    
    ret.Left = rightTree
    ret.Right = leftTree
    return ret
}
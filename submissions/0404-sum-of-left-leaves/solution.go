/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumOfLeft(root *TreeNode, isLeft bool) int {
    if root == nil {
        return 0
    }    
    if root.Left == nil && root.Right == nil && isLeft {
        return root.Val
    }
    if root.Left == nil && root.Right == nil {
        return 0
    }
    return sumOfLeft(root.Left, true) + sumOfLeft(root.Right, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
    return sumOfLeft(root, false)
}

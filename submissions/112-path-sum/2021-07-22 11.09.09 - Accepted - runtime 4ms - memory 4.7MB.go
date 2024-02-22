/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
      return false  
    } else if root.Left == nil && root.Right == nil && root.Val == targetSum {
        return true
    } else if root.Left == nil && root.Right == nil {
        return false
    } else if root.Left == nil  {
        return hasPathSum(root.Right, targetSum - root.Val)
    } else if root.Right == nil {
        return hasPathSum(root.Left, targetSum - root.Val) 
    } else {
        return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
    }
}
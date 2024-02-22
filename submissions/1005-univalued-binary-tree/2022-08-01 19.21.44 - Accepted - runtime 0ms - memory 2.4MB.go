/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    }    
    
    if root.Right != nil && root.Left !=nil{
        isUnival := isUnivalTree(root.Left) && isUnivalTree(root.Right)
        return isUnival && root.Val == root.Left.Val && root.Val == root.Right.Val
    }
    
    if root.Right != nil {
        return isUnivalTree(root.Right) && root.Val == root.Right.Val
    }
    
    if root.Left != nil {
        return isUnivalTree(root.Left) && root.Val == root.Left.Val
    }
    
    return true
}

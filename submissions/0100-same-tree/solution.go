/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil {
        return false
    } else if q == nil {
        return false
    } else if p.Val != q.Val {
        return false
    } else {
        var isSameLeft bool = isSameTree(p.Left, q.Left)
        var isSameRight bool = isSameTree(p.Right, q.Right)
        return isSameLeft && isSameRight
    }
}

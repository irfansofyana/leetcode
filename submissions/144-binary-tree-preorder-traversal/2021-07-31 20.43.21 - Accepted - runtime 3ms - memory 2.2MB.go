/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        tmp := make([]int, 0)
        return tmp
    }
    ans := make([]int, 0); ans = append(ans, root.Val)
    left := preorderTraversal(root.Left)
    right := preorderTraversal(root.Right)
    ans = append(ans, left...)
    ans = append(ans, right...)
    return ans
}
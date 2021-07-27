/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBSTHelper(root *TreeNode, minSoFar int64, maxSoFar int64) bool {
    if root == nil {
        return true
    }
    if int64(root.Val) <= minSoFar || int64(root.Val) >= maxSoFar {
        return false
    }
    return isValidBSTHelper(root.Left, minSoFar, int64(root.Val)) && 
    isValidBSTHelper(root.Right, int64(root.Val), maxSoFar)
}

func isValidBST(root *TreeNode) bool {
    return isValidBSTHelper(root, -10000000000, 10000000000)
}

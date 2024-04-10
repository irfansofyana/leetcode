/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    if root.Left == nil && root.Right == nil {
        return root.Val == 1
    }

    var valLeft bool = false
    if root.Left != nil {
        valLeft = evaluateTree(root.Left)
    }
    var valRight bool = false
    if root.Right != nil {
        valRight = evaluateTree(root.Right)
    }

    if root.Val == 2 {
        return valLeft || valRight
    }

    return valLeft && valRight
}

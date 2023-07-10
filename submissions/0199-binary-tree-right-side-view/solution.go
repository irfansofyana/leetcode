/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    view := make([]int, 0)
    left := make([]int, 0)
    right := make([]int, 0)

    view = append(view, root.Val)
    
    if root.Right != nil {
        right = rightSideView(root.Right)
    }
    if root.Left != nil {
        left = rightSideView(root.Left)
    }

    if len(right) >= len(left) {
        view = append(view, right...)
    } else {
        view = append(view, right...)
        view = append(view, left[len(right):]...)
    }

    return view
}

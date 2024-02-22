/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strconv"

func sumHelper(root *TreeNode, curr string) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        res, _ := strconv.Atoi(curr)
        return res
    }
    if root.Left == nil {
        return sumHelper(root.Right, curr + strconv.Itoa(root.Right.Val))
    }
    if root.Right == nil {
        return sumHelper(root.Left, curr + strconv.Itoa(root.Left.Val))
    }
    return sumHelper(root.Right, curr + strconv.Itoa(root.Right.Val)) +  sumHelper(root.Left, curr + strconv.Itoa(root.Left.Val))
}

func sumNumbers(root *TreeNode) int {
    return sumHelper(root, strconv.Itoa(root.Val))   
}
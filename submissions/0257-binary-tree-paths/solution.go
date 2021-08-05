/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"

func joinPath(ret []string, children []string, val int) []string {
    for _, child := range children {
        tmp := fmt.Sprintf("%d->%s", val, child)
        ret = append(ret, tmp)
    }
    return ret
}

func binaryTreePaths(root *TreeNode) []string {
    ret := make([]string, 0)
    if root == nil {
        return ret
    }
    if root.Left == nil && root.Right == nil {
        ret = append(ret, fmt.Sprintf("%d", root.Val))
        return ret
    }
    leftRet := binaryTreePaths(root.Left)
    rightRet := binaryTreePaths(root.Right)
    
    ret = joinPath(ret, leftRet, root.Val)
    ret = joinPath(ret, rightRet, root.Val)
    
    return ret
}

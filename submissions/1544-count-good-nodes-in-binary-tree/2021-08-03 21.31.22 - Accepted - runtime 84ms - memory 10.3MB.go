/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func cntGoodNodes(root *TreeNode, maks int) int {
    if root ==  nil {
        return 0
    }
    var cnt int = 0
    var maks2 int = maks
    if root.Val >= maks || maks == -100000000 {
        cnt += 1
    } 
    if root.Val > maks2 {
        maks2 = root.Val
    }
    return cnt + cntGoodNodes(root.Left, maks2) + cntGoodNodes(root.Right, maks2)
}

func goodNodes(root *TreeNode) int {
    return cntGoodNodes(root, -100000000)
}
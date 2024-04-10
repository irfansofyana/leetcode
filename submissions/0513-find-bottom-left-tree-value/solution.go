/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    ans := 0; maxDepth := -1
    dfs(root, 0, &ans, &maxDepth)
    return ans
}

func dfs(root *TreeNode, currDepth int, ans *int, maxDepth *int) {
    if root.Left == nil && root.Right == nil {
        if currDepth > *maxDepth {
            *maxDepth = currDepth
            *ans = root.Val
        }
    }

    if root.Left != nil {
        dfs(root.Left, currDepth + 1, ans, maxDepth)
    }
    if root.Right != nil {
        dfs(root.Right, currDepth + 1, ans, maxDepth)
    }
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var rootVal int = root.Val
    var qVal int = q.Val
    var pVal int = p.Val
    
    if pVal > qVal {
        pVal, qVal = qVal, pVal
    }
    if qVal > rootVal && pVal < rootVal {
        return root
    }
    if pVal == rootVal || qVal == rootVal {
        return root
    }
    if qVal < rootVal {
        return lowestCommonAncestor(root.Left, p, q)
    }
    return lowestCommonAncestor(root.Right, p, q)
}
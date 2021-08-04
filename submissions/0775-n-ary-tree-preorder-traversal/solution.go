/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    } 
    ans = append(ans, root.Val)
    
    for _, child := range root.Children {
        childRes := preorder(child)
        ans = append(ans, childRes...)
    }
    
    return ans
}

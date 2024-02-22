/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    ans := make([]int, 0)
    if root == nil {
        return ans
    } 

    for _, child := range root.Children {
        childRes := postorder(child)
        ans = append(ans, childRes...)
    }
    ans = append(ans, root.Val)
    
    return ans
}
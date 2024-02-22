/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    
    if len(root.Children) == 0 {
        return 1
    }

    maks := 0
    for _, child := range root.Children {
        depth := maxDepth(child)
        if depth > maks {
            maks = depth
        }
    }

    return maks+1
}
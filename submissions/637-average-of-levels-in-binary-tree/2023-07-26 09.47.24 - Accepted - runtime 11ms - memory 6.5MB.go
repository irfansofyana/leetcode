/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    sum := make(map[int]int)
    cnt := make(map[int]int)
    maxDepth := 0
    traverse(root, 0, &sum, &cnt, &maxDepth)
    
    avg := make([]float64, 0)
    for i := 0; i <= maxDepth; i++ {
        avg = append(avg, float64(sum[i]) / float64(cnt[i]) )
    }

    return avg
}

func traverse(root *TreeNode, level int, sum *map[int]int, cnt *map[int]int, maxDepth *int) {
    (*sum)[level] += root.Val
    (*cnt)[level]++

    if level > *maxDepth {
        *maxDepth = level
    }

    if root.Left != nil {
        traverse(root.Left, level+1, sum, cnt, maxDepth)
    }
    if root.Right != nil {
        traverse(root.Right, level+1, sum, cnt, maxDepth)
    }
}
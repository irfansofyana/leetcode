/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "math"

func maxLevelSum(root *TreeNode) int {
    sum := make(map[int]int)
    
    traverse(root, 1, &sum)
    
    minLevel := math.MaxInt64
    maxSum := math.MinInt64

    for level, sum := range sum {
        if sum > maxSum {
            maxSum = sum
            minLevel = level
        }
        if sum == maxSum && level < minLevel {
            minLevel = level
        }
    }

    return minLevel
}

func traverse(root *TreeNode, level int, sum *map[int]int) {
    (*sum)[level] += root.Val
    if root.Left != nil {
        traverse(root.Left, level+1, sum)
    }
    if root.Right != nil {
        traverse(root.Right, level+1, sum)
    }
}

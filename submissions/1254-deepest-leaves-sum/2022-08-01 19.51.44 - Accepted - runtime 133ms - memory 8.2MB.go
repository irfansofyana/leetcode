/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deepestLeavesSum(root *TreeNode) int {
    lvlTree := level(root)
    sum := 0
    sumLeaves(root, 1, lvlTree, &sum)
    return sum
}

func sumLeaves(root *TreeNode, currLvl int, maxLvl int, currSum *int) {
    if currLvl == maxLvl {
        *currSum += root.Val
        return
    }
    
    if root.Left != nil {
        sumLeaves(root.Left,currLvl+1,maxLvl,currSum)
    }
    
    if root.Right != nil {
        sumLeaves(root.Right,currLvl+1,maxLvl,currSum)
    }
}

func level(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    var lvl int = 0
    if root.Right != nil {
        rightLvl := level(root.Right)
        if rightLvl > lvl {
            lvl = rightLvl
        }
    }
    
    if root.Left != nil {
        leftLvl := level(root.Left)
        if leftLvl > lvl {
            lvl = leftLvl
        }
    }
    
    return lvl + 1
}
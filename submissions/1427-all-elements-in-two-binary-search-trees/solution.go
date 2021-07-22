/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getElements(root *TreeNode) []int{
    var elements = make([]int, 0)
    if root == nil {
        return elements
    } 
    var leftPart []int = getElements(root.Left)
    var rightPart []int = getElements(root.Right)
    
    elements = append(elements, leftPart...)
    elements = append(elements, root.Val)
    elements = append(elements, rightPart...)
    
    return elements
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    var firstList []int = getElements(root1)
    var secondList []int = getElements(root2)
    var ans = make([]int, 0)
    var iFirst int = 0
    var iSecond int = 0
    
    for iFirst < len(firstList) || iSecond < len(secondList) {
        if iFirst == len(firstList) {
            ans = append(ans, secondList[iSecond])
            iSecond += 1
        }else if iSecond == len(secondList) {
            ans = append(ans, firstList[iFirst])
            iFirst += 1
        } else if firstList[iFirst] < secondList[iSecond] {
            ans = append(ans, firstList[iFirst])
            iFirst += 1
        } else {
            ans = append(ans, secondList[iSecond])
            iSecond += 1
        }
    }
    
    return ans
}

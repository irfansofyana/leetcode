/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode
    var tmp = make([]int, 0)
    
    for {
        if l1 == nil && l2 == nil {
            break
        } else if l1 == nil {
            tmp = append(tmp, l2.Val)
            l2 = l2.Next
        } else if l2 == nil {
            tmp = append(tmp, l1.Val)
            l1 = l1.Next
        } else if l1.Val <= l2.Val {
            tmp = append(tmp, l1.Val)
            l1 = l1.Next
        } else {
            tmp = append(tmp, l2.Val)
            l2 = l2.Next
        }
    }
    
    if len(tmp) == 0 {
        return nil
    }
    
    result = &ListNode{tmp[len(tmp)-1], nil}
    for i := len(tmp)-2; i >= 0; i-- {
        result = &ListNode{tmp[i], result}
    }
    return result
}
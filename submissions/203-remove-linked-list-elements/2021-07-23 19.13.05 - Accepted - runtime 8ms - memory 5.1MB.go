/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var result *ListNode
    var prv *ListNode
    
    for head != nil {
        if head.Val != val && result == nil {
            result = &ListNode{head.Val, nil}
            prv = result
        } else if head.Val != val {
            prv.Next = &ListNode{head.Val, nil}
            prv = prv.Next
        } 
        head = head.Next
    }
    
    return result
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    reversedNext := reverseList(head.Next)
    ptr := reversedNext
    for ptr != nil {
        if ptr.Next == nil {
            ptr.Next = &ListNode{head.Val, nil}
            break
        }
        ptr = ptr.Next
    }
    return reversedNext
}
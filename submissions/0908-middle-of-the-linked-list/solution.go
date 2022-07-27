/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    arr := make([]*ListNode, 0)
    
    for head != nil {
        arr = append(arr, head)
        head = head.Next
    }
    
    totalEl := len(arr)
    
    return arr[totalEl/2]
}

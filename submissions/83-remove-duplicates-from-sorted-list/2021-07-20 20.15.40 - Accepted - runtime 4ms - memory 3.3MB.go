/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
    var ans *ListNode
    var currNext *ListNode
    var curr int = 0
    
    for head != nil {
        curr = head.Val
        if ans == nil {
            ans = &ListNode{curr, nil}
            currNext = ans
        } else {
            currNext.Next = &ListNode{curr, nil}
            currNext = currNext.Next
        }
        
        var pNext *ListNode = head.Next
        for pNext != nil && pNext.Val == curr {
            pNext = pNext.Next
        }
        head = pNext
    }
    
    return ans
}
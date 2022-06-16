/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    visited := make(map[*ListNode]bool)
    
    for head != nil {
        visited[head] = true
        
        if head.Next != nil && visited[head.Next] {
            return true
        }
        
        head = head.Next
    }
    
    return false
}

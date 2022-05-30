/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    //if node ahead is the same
    //replace head.Next with the 2nd removed node
    if head == nil {
        return nil
    }
    returnHead := head
    for head != nil {
        if head.Next == nil {
            break
        }
        if head.Val == head.Next.Val {
            hop := head.Next.Next
            head.Next = hop
            continue
        }
        head = head.Next
    }
    return returnHead
}
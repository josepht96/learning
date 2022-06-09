/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    //start at beginning of each list
    //add values
    //add to head of return list
    carryOver := 0
    var head *ListNode
    returnHead := head
    for l1 != nil || l2 != nil {
        var sum int
        if l1 != nil && l2 == nil {
            sum = l1.Val
            l1 = l1.Next
        }
        if l1 == nil && l2 != nil {
            sum = l2.Val
            l2 = l2.Next
        }
        if l1 != nil && l2 != nil {
            sum = l1.Val + l2.Val
            l1 = l1.Next
            l2 = l2.Next
        }
        sum = sum + carryOver
        if sum >= 10 {
            sum = sum - 10
            carryOver = 1
        } else {
            carryOver = 0
        }
        newNode := ListNode{sum, nil}
        if head == nil {
            head = &newNode
            returnHead = head
        } else {
            head.Next = &newNode
            head = head.Next
        }
    }
    if carryOver == 1 {
        head.Next = &ListNode{1, nil}
    }
    return returnHead
}

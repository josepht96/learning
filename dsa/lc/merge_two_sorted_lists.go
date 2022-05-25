func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var newList *ListNode
    head := newList
    if list1 == nil && list2 != nil {
        return list2
    } else if list2 == nil && list1 != nil {
        return list1
    } else if list1 == nil && list2 == nil {
        return head
    }
    if list1.Val <= list2.Val {
        newList = list1
        head = newList
        list1 =  list1.Next
    } else {
        newList = list2
        head = newList
        list2 = list2.Next
    }
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            newList.Next = list1
            newList = newList.Next
            list1 = list1.Next
        } else {
            newList.Next = list2
            newList = newList.Next
            list2 = list2.Next
        }
    }
    if list1 == nil {
        newList.Next = list2
    } else if list2 ==  nil {
        newList.Next = list1
    }
    return head
}
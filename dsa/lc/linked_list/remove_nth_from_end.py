# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        trailing_ptr = head
        forward_ptr = head
        if head.next is None or head is None:
            return None
        # stay 1 behind replacement node
        i = 0
        while i != n:
            forward_ptr = forward_ptr.next
            i += 1
        if forward_ptr is not None:
            while forward_ptr.next is not None:
                forward_ptr = forward_ptr.next
                trailing_ptr = trailing_ptr.next
        else:
            # if forward_ptr is None after going lenght of n,
            # we know that n == length of the list
            return head.next
            
        trailing_ptr.next = trailing_ptr.next.next
        return head
        
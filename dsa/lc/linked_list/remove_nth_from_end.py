# Definition for singly-linked list.
# dont think I did this right
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
        
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        step, count = 1, 1
        return_head = head

        while head.next != None:
            head = head.next
            count += 1
        head = return_head
        break_node = count - n
        if break_node == 0:
            return head.next
        while step < break_node:
            head = head.next
            step += 1
        if head.next.next == None:
            head.next = None
        else:
            temp = head.next.next
            head.next = temp

        return return_head
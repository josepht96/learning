# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # new head is going to be the original tail of the LL
        # think about how you start and end...
        next_node = None
        prev_node = None
        while head != None:
            if head.next == None:
                head.next = prev_node
                return head
            next_node = head.next
            head.next = prev_node
            prev_node = head
            head = next_node


        return head
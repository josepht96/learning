# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        tail = head
        count = 1
        step = 1
        mid = 0
        while tail.next != None:
            tail = tail.next
            count += 1
        mid = (count // 2) + 1
        while step < mid:
            head = head.next
            step += 1

        return head
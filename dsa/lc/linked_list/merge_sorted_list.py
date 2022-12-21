# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        # if one list ends, add the other to the end
        if list1 == None and list2 == None:
            return None
        elif list1 == None and list2 != None:
            return list2
        elif list1 != None and list2 == None:
            return list1

        return_list = None
        if list1.val <= list2.val:
            return_list = list1
            list1 = list1.next
        else:
            return_list = list2
            list2 = list2.next

        return_head = return_list

        while list1 != None and list2 != None:
            if list1.val <= list2.val:
                return_list.next = list1
                list1 = list1.next
            else:
                return_list.next = list2
                list2 = list2.next

            return_list = return_list.next

        if list1 == None and list2 != None:
            return_list.next = list2
        elif list1 != None and list2 == None:
            return_list.next = list1
        
        return return_head
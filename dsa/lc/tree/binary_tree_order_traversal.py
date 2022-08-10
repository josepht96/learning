# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return
        response = []
        queue = [root]
        
        next_level_length = 0
        current_level_length = 1
        current_array = []
        
        while len(queue) > 0:
            current_node = queue.pop(0)
            
            if current_node.left is not None:
                next_level_length += 1
                queue.append(current_node.left)
                
            if current_node.right is not None:
                next_level_length += 1
                queue.append(current_node.right)
                
            current_array.append(current_node.val)
            if len(current_array) == current_level_length:
                response.append(current_array)
                current_array = []
                current_level_length = next_level_length
                next_level_length = 0
            
        return response
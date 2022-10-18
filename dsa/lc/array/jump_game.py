class Solution:
    def canJump(self, nums: List[int]) -> bool:
        # if each index can reach the closest_good_index
        # then it is also a good index, and thus becomes closest_good_index
        closest_good_index = len(nums) - 1
        for i in range(len(nums),0,-1):
            j = i - 1
            if j + nums[j] >= closest_good_index:
                closest_good_index = j
        
        if closest_good_index == 0:
            return True
        else:
            return False
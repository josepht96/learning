class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l_ptr = 0
        r_ptr = len(nums) - 1
        while l_ptr <= r_ptr:
            if nums[l_ptr] == 0:
                nums.pop(l_ptr)
                nums.append(0)
                r_ptr -= 1
            else:
                l_ptr += 1
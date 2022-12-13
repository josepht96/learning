class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        If you can return an array...
        """
        length = len(nums)-1
        newArr = nums[length-k+1:] + nums[:k+1]
        nums = newArr

class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i = 0
        length = len(nums) - 1
        while i < k:
            last = nums[length]
            nums.insert(0, last)
            nums.pop(length+1)
            i +=1
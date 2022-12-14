class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        # key is the diff, value is the index
        diff_map = {}
        for i, num in enumerate(numbers):
            diff = target - num
            if diff in diff_map:
                return [diff_map[diff]+1, i+1]
            else:
                diff_map[num] = i
        
        return []
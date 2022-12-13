class TimSort:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        returnArr = []
        for i, num in enumerate(nums):
            sqr = num * num
            returnArr.append(sqr)
            i+=1

        returnArr.sort()
        return returnArr

class TwoPointer:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        returnArr = []
        l, r = 0, len(nums)-1

        while l <= r:
            if abs(nums[r]) >= abs(nums[l]):
                returnArr.append(nums[r] ** 2)
                r -= 1
            else:
                returnArr.append(nums[l] ** 2)
                l +=1

        return reversed(returnArr)
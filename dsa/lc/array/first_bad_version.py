# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        # start at left, start at right
        l = 0
        r = n
        if n == 1:
            return 1
        while l <= r:
            mid = (l+r)//2
            if not isBadVersion(mid):
                l = mid + 1
            elif isBadVersion(mid-1) != True:
                return mid
            else:
                r = mid - 1
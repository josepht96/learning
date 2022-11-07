class Solution:
    def maximum69Number (self, num: int) -> int:
        factor = int(math.log10(num))
        max = num
        while factor >= 0:
            digit = self.get_digit(num, factor)
            if digit == 6:
                newNum = 3 * (10**(factor)) + num
                if newNum > max:
                    max = newNum
            factor -=1
        return max
                
    def get_digit(self, number, n):
        return number // 10**n % 10
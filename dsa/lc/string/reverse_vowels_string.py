class Solution:
    def reverseVowels(self, s: str) -> str:
        bank = {'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
        i = 0
        j = len(s) - 1
        returnStr = ''
        for i in range(len(s)):
            if s[i] not in bank:
                returnStr += s[i]
            else:
                while s[j] not in bank and j != -1:
                    j -= 1
                returnStr += s[j]
                j -= 1
                
        return returnStr
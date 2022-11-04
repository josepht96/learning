class Solution:
    def longestPalindrome(self, words: List[str]) -> int:
        # zero or one middle duplicate pair
        # any additional duplicates must have a matching pair
        count = 0
        center = False
        hashMap = {}
        for word in words:
            if word not in hashMap:
                hashMap[word] = 1
            else:
                val = hashMap[word]
                hashMap[word] = val + 1
                
        for key in hashMap:
            revKey = key[1] + key[0]
            if revKey in hashMap:
                if key == revKey:
                    occurrences = hashMap[key]
                    if occurrences % 2 == 0:
                        count += occurrences
                    else:
                        center = True
                        count += occurrences - 1
                else:
                    occurrences = min(hashMap[key], hashMap[revKey])
                    count += occurrences
        if center:
            count += 1
            
        return count * 2
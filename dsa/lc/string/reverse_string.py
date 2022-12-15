class Solution:
    def reverseString(self, s: List[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        l_ptr = 0
        r_ptr = len(s) - 1
        while l_ptr <= r_ptr:
            l_char = s[l_ptr]
            r_char = s[r_ptr]
            s[l_ptr] = r_char
            s[r_ptr] = l_char
            l_ptr += 1
            r_ptr -= 1
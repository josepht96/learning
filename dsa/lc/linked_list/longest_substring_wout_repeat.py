class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        # sliding window
        l_ptr, r_ptr = 0, 0
        longest_window = 0
        char_map = {}
        while r_ptr <= len(s) - 1:
            if s[r_ptr] not in char_map:
                char_map[s[r_ptr]] = True
                r_ptr += 1
            else:
                char_map.pop(s[l_ptr])
                l_ptr += 1

            if r_ptr - l_ptr > longest_window:
                longest_window = r_ptr - l_ptr

        return longest_window
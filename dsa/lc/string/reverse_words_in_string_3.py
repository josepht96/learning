class Solution:
    def reverseWords(self, s: str) -> str:
        l_ptr = 0
        r_ptr = 0
        reverse_string = ""
        end_index = len(s) - 1
        while r_ptr <= end_index:
            if s[r_ptr] == " ":
                jump_index = r_ptr + 1
                r_ptr -= 1
                while l_ptr <= r_ptr:
                    reverse_string += s[r_ptr]
                    r_ptr -= 1
                l_ptr = jump_index
                r_ptr = jump_index
                reverse_string += " "

            elif r_ptr == end_index:
                while l_ptr <= r_ptr:
                    reverse_string += s[r_ptr]
                    r_ptr -= 1
                return reverse_string

            else:
                r_ptr += 1
        return reverse_string

    def reverseWord(self, w: str) -> str:
            l_ptr = 0
            r_ptr = len(w) - 1
            rev_word = ""
            while l_ptr <= r_ptr:
                l_char = w[l_ptr]
                r_char = w[r_ptr]
                

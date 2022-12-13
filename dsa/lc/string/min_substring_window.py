class Solution:
    def minWindow(self, s: str, t: str) -> str:
        # need way to check if string works or not
        # when a valid block is found, reduce left side until equal to right side
        # check if its valid shortest block
        l_ptr = 0
        r_ptr = len(t)-1
        length = len(s)-1
        t_length = len(t)-1
        char_map = {}
        min_window =  ""
        if s == t:
            return s
        # fine
        for char in t:
            if char in char_map:
                char_map[char] +=1
            else:
                char_map[char] = 1
        while l_ptr <= r_ptr:
            window = s[l_ptr:r_ptr+1]
            valid = self.validateWindow(char_map, window)
            if valid:
                if min_window == "":
                    min_window = window
                elif len(window) < len(min_window):
                   min_window = window
                l_ptr +=1
            elif r_ptr == length:
                break
            else:
                r_ptr +=1
        return min_window

    def validateWindow(self, char_map: map, window: str) -> bool:
        temp_map = char_map.copy()
        for char in window:
            if char in temp_map:
                if temp_map[char] == 1:
                    temp_map.pop(char, None)
                else:
                    temp_map[char] -=1
        if len(temp_map) != 0:
                return False

        return True

class Solution2:
    def minWindow(self, s: str, t: str) -> str:
        # need way to check if string works or not
        # when a valid block is found, reduce left side until equal to right side
        # check if its valid shortest block
        l_ptr = 0
        r_ptr = 0
        char_map = {}
        min_window =  ""
        if s == t:
            return s
        char_map = collections.Counter(t)
        while r_ptr < len(s):
            if s[r_ptr] in char_map:
                char_map[s[r_ptr]] -= 1

            valid = self.validateMap(char_map)
            window = s[l_ptr:r_ptr+1]
            if valid:
                while l_ptr <= r_ptr:
                    window = s[l_ptr:r_ptr+1]
                    if min_window == "":
                        min_window = window
                    elif len(min_window) > len(window):
                        min_window = window
                    
                    if s[l_ptr] in char_map:
                        char_map[s[l_ptr]] +=1
                    l_ptr += 1
                    valid = self.validateMap(char_map)
                    if not valid:
                        r_ptr +=1
                        break

            elif not valid:
                r_ptr += 1      

        return min_window

    def validateMap(self, char_map: map) -> bool:
        for key in char_map:
            if char_map[key] > 0:
                return False
        return True
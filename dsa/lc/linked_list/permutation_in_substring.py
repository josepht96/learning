class Solution1:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        # sliding window
        char_map = {}
        char_map = collections.Counter(s1)
        l_ptr, r_ptr = 0, len(s1)-1
        while r_ptr <= len(s2)-1:
            current_map = {}
            current_map = collections.Counter(s2[l_ptr:r_ptr+1])
            if current_map == char_map:
                return True
            else:
                l_ptr += 1
                r_ptr += 1
        return False

class Solution2:
     def checkInclusion(self, s1: str, s2: str) -> bool:
        # sliding window
        char_map = {}
        char_map = collections.Counter(s1)
        l_ptr, r_ptr = 0, len(s1)-1
        sliding_map = collections.Counter(s2[l_ptr:r_ptr+1])
        while r_ptr <= len(s2)-1:
            if sliding_map == char_map:
                return True
            else:
                # handle change to sliding_map
                if sliding_map[s2[l_ptr]] == 1:
                    sliding_map.pop(s2[l_ptr], None)
                else:
                    sliding_map[s2[l_ptr]] -= 1
                
                l_ptr += 1
                r_ptr += 1
                if r_ptr <= len(s2) - 1:
                    if s2[r_ptr] not in sliding_map:
                        sliding_map[s2[r_ptr]] = 1
                    else:
                        sliding_map[s2[r_ptr]] += 1
        return False
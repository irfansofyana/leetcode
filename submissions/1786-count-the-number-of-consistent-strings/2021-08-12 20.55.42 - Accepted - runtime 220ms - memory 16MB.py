class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        cnt = 0
        for word in words:
            isValid = True
            for ch in word:
                if ch not in allowed:
                    isValid = False
                    break
            if isValid:
                cnt += 1
        return cnt
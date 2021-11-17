class Solution(object):
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        vowels = []
        for i in range(len(s)):
            if s[i].lower() in ['a', 'i', 'u', 'e', 'o']:
                vowels.append(s[i])
                
        res = ""
        idx = len(vowels) - 1
        for i in range(len(s)):
            if s[i].lower() in ['a', 'i', 'u', 'e', 'o']:
                res = res + vowels[idx]
                idx -= 1
            else:
                res = res + s[i]
                
        return res
        

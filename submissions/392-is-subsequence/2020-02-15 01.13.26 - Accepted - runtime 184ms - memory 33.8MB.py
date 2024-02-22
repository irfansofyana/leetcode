class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        indeks_s = 0
        n = len(t)
        if (s == ""):
            return True
        for i in range(0, n):
            if (t[i] == s[indeks_s]):
                indeks_s += 1
            if (indeks_s == len(s)):
                return True
        return False
        
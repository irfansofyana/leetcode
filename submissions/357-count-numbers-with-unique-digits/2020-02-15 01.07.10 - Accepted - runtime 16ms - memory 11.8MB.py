class Solution(object):
    def countNumbersWithUniqueDigits(self, n):
        """
        :type n: int
        :rtype: int
        """
        ans = 0
        for i in range(1, min(n+1, 11)):
            if (i == 1):
                ans += 10
            else:
                tmp = 9
                kali = 9
                for j in range(1, i):
                    kali *= tmp
                    tmp -= 1
                ans += kali
        if (n == 0):
            ans += 1
        return ans
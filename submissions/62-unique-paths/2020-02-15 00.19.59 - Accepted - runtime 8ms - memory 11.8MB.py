class Solution(object):
    def uniquePaths(self, m, n):
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        dp = []
        for i in range(m):
            cell = []
            for j in range(n):
                cell.append(0)
            dp.append(cell)
        for i in range(1, n):
            dp[0][i] += 1
        for i in range(1, m):
            dp[i][0] += 1
        for i in range(1, m):
            for j in range(1, n):
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
        if (n == 1 and m == 1):
            return 1
        else:
            return dp[m-1][n-1]
        
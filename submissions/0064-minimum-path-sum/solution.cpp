class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) {
        vector<vector<int> > dp;
        int n = (int)grid.size(); int m = (int)grid[0].size();
        for (int i = 0; i < n; ++i)
            dp.push_back(grid[i]);
        for (int i = 0; i < m; ++i)
            dp[0][i] = (i == 0 ? 0 : dp[0][i-1]) + grid[0][i];
        for (int i = 0; i < n; ++i)
            dp[i][0] = (i == 0 ? 0 : dp[i-1][0]) + grid[i][0];
        for (int i = 1; i < n; ++i){
            for (int j = 1; j < m; ++j) {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j];
            }
        }
        
        return dp[n-1][m-1];
    }
};

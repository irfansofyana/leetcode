class Solution {
public:
    int dp[71][71][71];
    int rows, cols;
    vector<vector<int>> cgrid;
    
    bool isSame(int c1, int c2) {
        return (c1 == c2);
    }
    
    bool isValidCol(int c) {
        return (c >= 0 && c < cols);
    }
    
    int findMax(int row, int c1, int c2) {
        if (row == rows) {
            return 0;
        }
        
        if (dp[row][c1][c2] != -1) {
            return dp[row][c1][c2];
        }
        
        int sum = cgrid[row][c1];
        if (!isSame(c1, c2)) {
            sum += cgrid[row][c2];
        }
        
        int maxx = -1e9;
        for (int i = -1; i <= 1; i++) {
            for (int j = -1; j <= 1; j++) {
                int nC1 = c1 + i;
                int nC2 = c2 + j;
                if (!isValidCol(nC1) || !isValidCol(nC2)) {
                    continue;
                }
                maxx = max(sum + findMax(row+1, nC1, nC2), maxx);
            }
        }
        
        return dp[row][c1][c2] = maxx;
    }
    
    int cherryPickup(vector<vector<int>>& grid) {
        memset(dp, -1, sizeof dp);
        cgrid = grid;
        rows = (int)grid.size();
        cols = (int)grid[0].size();
        return findMax(0, 0, cols-1);
    }
};
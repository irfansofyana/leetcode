class Solution {
public:
    int dp[10000005];
    
    int numSquares(int n) {
        dp[0] = 1;
        for (int i = 1; i <= n; ++i){
            int k = (int)sqrt(i);
            int mini=-1;
            if (k * k == i){
                mini = 1;
            }else {
                for (int j = 1; j <= k; ++j){
                    mini = (j == 1 ? dp[i-j*j]+1:min(mini, dp[i-j*j]+1));
                }
            }
            dp[i] = mini;
        }
        return dp[n];
    }
};
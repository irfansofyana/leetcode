class Solution {
public:
    int dp[1005][35], N, K;

    const int MOD = 1000000007;

    int numRollsToTarget(int n, int k, int target) {
        memset(dp, -1, sizeof dp);
        N = n; K = k;
        return findCount(target, 0);    
    }

    int findCount(int target, int curr) {
        if (curr == N && target == 0) {
            return 1;
        } 
        if (curr == N) {
            return 0;
        }
        if (dp[target][curr] != -1) {
            return dp[target][curr];
        }

        int tot = 0;
        for (int i = 1; i <= min(K, target); i++) {
            tot += findCount(target-i, curr+1);
            tot %= MOD;
        }

        return dp[target][curr] = tot;
    }
};

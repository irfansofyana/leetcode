class Solution {
public:
    int maks[303][303];
    int dp[303][13], n;
        
    int cari(int idx, int sisa){
        if (sisa == 1){
            return maks[idx][n-1];
        }
        int &ret = dp[idx][sisa];
        if (ret != -1) return ret;
        ret = 1000000000;
        for (int i = idx; i <= n-sisa; ++i){
            ret = min(ret, maks[idx][i] + cari(i+1, sisa-1));
        }
        return ret;
    }
    int minDifficulty(vector<int>& jobDifficulty, int d) {
        if (d > jobDifficulty.size()) return -1;
        memset(dp, -1, sizeof dp);
        n = jobDifficulty.size();
        for (int i = 0; i < n; ++i){
            for (int j = i; j < n; ++j){
                if (i == j) maks[i][j] = jobDifficulty[i];
                else maks[i][j] = max(maks[i][j-1], jobDifficulty[j]);
            }
        }
        int ans = cari(0, d);
        return (ans == 1000000000 ? -1 : ans);
    }
};
class Solution {
public:
    vector<int> A;
    int dp[501][501], maks[501][501];
    int n, limit;
    
    
    void generateMaksArray() {
        for (int i = 0; i < n; i++) {
            int _maks = A[i];
            for (int j = i; j < n; j++) {
                _maks = max(_maks, A[j]);
                maks[i][j] = _maks;
            }
        }    
    }
    
    int findOptimumPartition(int now, int st) {
        if (st == n) {
            return 0;
        }
        if (now == n) {
            return maks[st][now-1] * (now-st);
        }
        if (dp[now][st] != -1) {
            return dp[now][st];
        }
        
        
        int optimum = maks[st][now] * (now-st+1) + findOptimumPartition(now + 1, now + 1);
        if (now-st+1 != limit) {
            optimum = max(optimum, findOptimumPartition(now + 1, st));   
        } 
        
        return dp[now][st] = optimum;
    }
        
    int maxSumAfterPartitioning(vector<int>& arr, int k) {
        n = (int)arr.size(); limit = k;
        for (int i = 0; i < n; i++) {
            A.push_back(arr[i]);
        }
        
        memset(dp, -1, sizeof dp);
        generateMaksArray();
        
        return findOptimumPartition(0, 0);
    }
};

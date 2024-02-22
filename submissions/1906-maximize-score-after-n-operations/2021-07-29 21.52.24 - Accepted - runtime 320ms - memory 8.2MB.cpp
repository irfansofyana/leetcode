class Solution {
public:
    int dp[8][(1<<14) + 2];
    vector<int> arr;
    int n;
    
    int findMax(int step, int state) {
        if (step > n) return 0;
        if (dp[step][state] != -1) return dp[step][state];
        int maks =-1e9;
        int m = 2 * n;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < m; ++j) {
                if ((state&(1<<i)) == 0 && (state&(1<<j)) == 0 && i != j) {
                    int nextState = (state|(1<<i));
                    nextState |= (1<<j);
                    maks = max(maks, step * __gcd(arr[i], arr[j]) + findMax(step+1, nextState));
                }
            }
        }
        return dp[step][state] = maks;
    }
    
    int maxScore(vector<int>& nums) {
        memset(dp, -1, sizeof dp);
        for (int i = 0; i < nums.size(); ++i) {
            arr.push_back(nums[i]);
        }
        n = (int)nums.size() / 2;
        
        return findMax(1, 0);
    }
};
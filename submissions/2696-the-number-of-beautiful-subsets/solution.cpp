class Solution {
public:
    vector<int> arr;
    int dp[(1<<20)+2][21], N, K;


    int beautifulSubsets(vector<int>& nums, int k) {
        memset(dp, -1, sizeof dp); 
        N = nums.size(); K = k;
        arr = nums;   
        return countWays(0, 0);
    }

    int countWays(int bitmask, int pos) {
        if (pos == N && bitmask > 0) {
            return 1;
        }
        if (pos == N) {
            return 0;
        }

        if (dp[bitmask][pos] != -1) {
            return dp[bitmask][pos];
        }

        bool isExist = false;
        for (int i = 0; i < N && !isExist; i++) {
            if ((bitmask & (1<<i)) > 0) {
                if (abs(arr[i] - arr[pos]) == K) {
                    isExist = true;
                }
            }
        }
        
        // not take
        int total = countWays(bitmask, pos+1);
        if (!isExist) {
            // take because possible
            total += countWays((bitmask|(1 << pos)), pos+1);
        }

        return dp[bitmask][pos] = total;
    }
};

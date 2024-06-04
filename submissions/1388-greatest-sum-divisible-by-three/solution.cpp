class Solution {
public:
    vector<int> arr;
    int dp[40005][4], N;

    int maxSumDivThree(vector<int>& nums) {
        N = nums.size(); arr = nums;
        memset(dp, -1, sizeof dp);
        return findMax(0, 0);
    }

    int findMax(int pos, int rem) {
        if (pos == N && rem == 0) {
            return 0;
        }
        if (pos == N) {
            return -1000000007;
        }
        if (dp[pos][rem] != -1) {
            return dp[pos][rem];
        }

        int curr_max = findMax(pos+1, (rem-(arr[pos]%3)+3)%3) + arr[pos];
        curr_max = max(findMax(pos+1, rem), curr_max);
        
        return dp[pos][rem] = curr_max;
    }
};  

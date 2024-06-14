class Solution {
public:
    long long dp[100003][3];
    vector<int> nums;

    long long maxAlternatingSum(vector<int>& nums) {
        memset(dp, -1, sizeof dp);
        this->nums = nums;
        return findMax(0, 0);
    }

    long long findMax(int pos, int parity) {
        if (pos == nums.size()) {
            return 0;
        }

        if (dp[pos][parity] != -1) {
            return dp[pos][parity];
        }

        long long maks = findMax(pos+1, parity);
        long long val;
        if (parity % 2 == 0) {
            val = findMax(pos+1, (parity+1)%2) + 1LL*nums[pos];
        } else {
            val = findMax(pos+1, (parity+1)%2) - 1LL*nums[pos];
        }
        maks = max(maks, val);

        return dp[pos][parity] = maks;
    }
};

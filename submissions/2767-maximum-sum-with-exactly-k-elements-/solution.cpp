class Solution {
public:
    int maximizeSum(vector<int>& nums, int k) {
        int maks = -1;
        for (int i = 0; i < nums.size(); i++) {
            maks = max(maks, nums[i]);
        }
        int ans = 0;
        for (int i = 0; i < k; i++) {
            ans += maks;
            maks++;
        }
        return ans;
    }
};

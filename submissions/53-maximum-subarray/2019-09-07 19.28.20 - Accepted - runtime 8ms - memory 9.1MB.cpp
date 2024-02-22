class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int res = 0;
        int ans = -1;
        for (int i = 0; i < (int)(nums.size()); ++i){
            res += nums[i];
            if (i == 0) ans = res;
            else ans = max(res, ans);
            if (res < 0) res = 0;
        }
        return ans;
    }
};
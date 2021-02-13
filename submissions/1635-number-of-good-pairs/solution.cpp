class Solution {
public:
    int numIdenticalPairs(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int i = 0;
        int ans = 0;
        while (i < nums.size()) {
            int j = i + 1;
            int curr = 1;
            while (j < nums.size() && nums[i] == nums[j]) {
                ++j;
                ++curr;
            }
            ans += curr * (curr - 1) / 2;
            i = j;
        }
        
        return ans;
    }
};

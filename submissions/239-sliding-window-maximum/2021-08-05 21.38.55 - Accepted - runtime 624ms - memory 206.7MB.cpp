class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        vector<int> ans;
        multiset<int> ms;
        int n = (int)nums.size();
        
        for (int i = 0; i < n; i++) {
            ms.insert(nums[i]);
            if (i >= k-1) {
                auto it = ms.end(); it--;
                ans.push_back(*it);
                ms.erase(ms.find(nums[i-k+1]));
            }
        }
        
        return ans;
    }
};
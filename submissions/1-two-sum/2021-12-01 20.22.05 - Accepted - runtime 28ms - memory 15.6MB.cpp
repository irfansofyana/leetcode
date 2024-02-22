class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, vector<int> > arr;
        for (int i = 0; i < (int)nums.size(); ++i) {
            arr[nums[i]].push_back(i);
        }
        vector<int> ans;
        for (int i = 0; i < (int)nums.size(); ++i){
            int other = target - nums[i];
            if (other != nums[i] && arr[other].size() > 0) {
                ans.push_back(i); ans.push_back(arr[other][0]);
                break;
            } else if (other == nums[i] && arr[other].size() > 1) {
                for (int j = 0; j < 2; ++j) ans.push_back(arr[other][j]);
                break;
            }
        }
        return ans;
    }
};
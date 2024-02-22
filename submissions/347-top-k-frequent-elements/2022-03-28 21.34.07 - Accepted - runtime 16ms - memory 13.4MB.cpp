class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        int n = (int)nums.size();
        sort(nums.begin(), nums.end());
        priority_queue<pair<int,int> > pq;
        
        int i = 0;
        while (i < n) {
            int j = i + 1;
            int freq = 1;
            while (j < n && nums[j] == nums[i]) {
                freq++;
                j++;
            }
            pq.push(make_pair(freq, nums[i]));
            i = j;
        }
        
        vector<int> ans;
        for (int i = 0; i < k; i++) {
            pair<int,int> p = pq.top();
            pq.pop();
            ans.push_back(p.second);
        }
        
        return ans;
    }
};
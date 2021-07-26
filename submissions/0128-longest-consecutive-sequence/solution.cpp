class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        map<int, bool> isExist;
        int n = (int)nums.size();
        int ans = 0, curr = 0, streak=0;
        
        for (int i = 0; i < n; ++i) {
            isExist[nums[i]] = true;
        }
        
        auto it = isExist.begin();
        while (it != isExist.end()) {
            if (it == isExist.begin()) {
                curr = it->first;
                streak = 1;
            } else {
                if (it->first - curr == 1) {
                    curr = it->first;
                    streak++;
                } else {
                    curr = it->first;
                    streak = 1;
                }
            }
            ans = max(ans, streak);
            ++it;
        }
        
        return ans;
    }
};

class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        map<int,vector<int> > mep;
        int sum = 0;
        for (int i = 0; i < (int)nums.size(); ++i){
            sum += nums[i];
            mep[sum].push_back(i);
        }
        int ans = 0;
        int cur = 0;
        for (int i = 0; i < (int)nums.size(); ++i){
            int target = k + cur;
            int lo = 0;
            int sz = (int)mep[target].size();
            int hi = sz - 1;
            if (lo <= hi){
                int pos = -1;
                while (lo <= hi){
                    int mid = (lo + hi) >> 1;
                    if (mep[target][mid] >= i){
                        pos = mid;
                        hi = mid-1;
                    }else {
                        lo = mid+1;
                    }
                } 
                if (pos != -1){
                    ans += (sz-pos);
                }
            }
            cur += nums[i];
        }
        return ans;
    }
};
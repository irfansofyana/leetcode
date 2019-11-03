class Solution {
public:
    int maxEqualFreq(vector<int>& nums) {
        int n = nums.size();
        map<int,int> frek;
        map<int,int> distinct;
        int ans = 0;
        for (int i = 0; i < n; ++i){
            int t = frek[nums[i]];
            if (t == 0){
                frek[nums[i]]++;
                distinct[1]++;
            }else {
                distinct[t]--;
                if (distinct[t] == 0) {
                    distinct.erase(t);
                }
                frek[nums[i]]++;
                distinct[t+1]++;
            }
            int sz = distinct.size();
            if (sz == 1){
                auto be = distinct.begin();
                if (be->first == 1) ans = max(ans, i+1);
                else if (be->second == 1) ans = max(ans, i+1);
            }else if (sz == 2){
                auto en = distinct.end();
                --en;
                auto be = distinct.begin();
                if (be->second == 1 && be->first == 1){
                    ans = max(ans, i+1);
                }
                if (en->second == 1 && en->first == be->first + 1){
                    ans = max(ans, i+1);
                }
            }
        }
        return ans;
    }
};

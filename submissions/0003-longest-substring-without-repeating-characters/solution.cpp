class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n = (int)s.size();
        int l = 0;
        int r = 0;
        int frek[1005]={0};
        int ans = 0;
        while (r < n){
            int idx = s[r];
            if (frek[idx] == 0){
                ++frek[idx];
                ans = max(ans, r-l+1);
            }else {
                while (l <= r && frek[idx] > 0){
                    --frek[s[l]];
                    ++l;
                }
                ++frek[idx];
                ans = max(ans, r-l+1);
            }
            ++r;
        }
        return ans;
    }
};


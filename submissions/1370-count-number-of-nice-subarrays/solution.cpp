class Solution {
public:
    int numberOfSubarrays(vector<int>& nums, int k) {
        int n = (int)nums.size();
        int dp[n+1] = {0};
        int ans = 0;
        for (int i = 0; i < n; ++i){
            dp[i] = (i == 0 ? (nums[i]%2 == 1) : dp[i-1]+(nums[i]%2 == 1));
        }
        for (int i = 0; i < n; ++i){
            int pre = (i == 0 ? 0 : dp[i-1]);
            int lo = i, hi = n-1;
            int ll = -1, rr = -1;
            while (lo <= hi){
                int mid = (lo + hi) >> 1;
                if (dp[mid] < k+pre){
                    lo = mid+1;
                }else if (dp[mid] > k+pre){
                    hi = mid-1;
                }else {
                    ll = (ll == -1 ? mid : min(ll, mid));
                    hi = mid-1;
                }
            }
            rr = ll;
            lo = i, hi = n-1;
            while (lo <= hi){
                int mid = (lo + hi) >> 1;
                if (dp[mid] < k+pre){
                    lo = mid+1;
                }else if (dp[mid] > k+pre){
                    hi = mid-1;
                }else {
                    rr = (rr == -1 ? mid : max(rr, mid));
                    lo = mid+1;
                }
            }
            if (ll == -1 || rr == -1) continue;
            ans += rr-ll+1;
        }
        return ans;
        //1 2 2 3 4
    }
};

class Solution {
public:
    long long dp[1003][55];
    long long prefiks[1005];
    int n;
    
    long long query(int a,int b){
        return prefiks[b] - (a == 0 ? 0 : prefiks[a-1]);    
    }
    
    long long cari(int pos, int rem){
        if (pos == n && rem == 0){
            return 0;
        }
        if (pos == n || rem < 0){
            return 1e18;
        }
        if (rem == 0 && pos != n){
            return 1e18;
        }
        long long &ret = dp[pos][rem];
        if (ret != -1) return ret;
        for (int i = pos; i < n; ++i){
            if (ret == -1){
                ret = max(query(pos, i), cari(i+1, rem-1));
            }else {
                ret = min(ret, max(query(pos, i), cari(i+1, rem-1)));
            }
        }
        return ret;
    }
    
    int splitArray(vector<int>& nums, int m) {
        n = (int)nums.size();
        for (int i = 0; i < n; ++i){
            prefiks[i] = (i == 0 ? nums[i] : nums[i]+prefiks[i-1]);
        }
        memset(dp, -1, sizeof dp);
        return cari(0, m);
    }
};

class Solution {
public:
    int dp[505][505];
    int n;
    vector<int> arr;
    
    int cari(int pos, int _time) {
        if (pos == n) {
            return 0;
        }
        if (dp[pos][_time] != -1) {
            return dp[pos][_time];
        }
        
        int ambil = arr[pos] * _time + cari(pos + 1, _time + 1);
        int notAmbil = cari(pos + 1, _time);
        
        return dp[pos][_time] = max(ambil, notAmbil);
    }
    
    int maxSatisfaction(vector<int>& satisfaction) {
        n = (int)satisfaction.size();
        for (int i = 0; i < n; ++i) {
            arr.push_back(satisfaction[i]);
        }
        
        sort(arr.begin(), arr.end());
        
        memset(dp, -1, sizeof dp); 
        
        int ans = cari(0, 1);
        if (ans < 0) ans = 0;
        
        return ans;
    }
};
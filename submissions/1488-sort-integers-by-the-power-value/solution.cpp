class Solution {
public:
    map<int,int> dp;

    int getKth(int lo, int hi, int k) {
        for (int i = 1; i <= 1000; i++) {
            dp[i] = getVal(i);
        }
        vector<pair<int, int>> v;
        for (int i = lo; i <= hi; i++) {
            v.push_back({dp[i], i});
        }
        sort(v.begin(), v.end());

        return v[k-1].second;    
    }

    int getVal(int n) {
        if (n == 1) {
            return 0;
        }
        if (dp.find(n) != dp.end()) {
            return dp[n];
        }
        if (n % 2 == 0) {
            return dp[n] = getVal(n/2) + 1;
        }

        return dp[n] = getVal(3*n+1) + 1;
    }
};

class Solution {
public:
    int dp[105][205], destination;
    vector<int> locs;
    const int MOD = 1000000007;

    int countRoutes(vector<int>& locations, int start, int finish, int fuel) {
        locs = locations; destination = finish;
        memset(dp, -1, sizeof dp);

        return ways(start, fuel);
    }

    int ways(int source, int currFuel) {
        if (source == destination) {
            int ans = 1;
            for (int i = 0; i < locs.size(); i++) {
                int used = abs(locs[i]-locs[source]);
                if (i != source &&  used <= currFuel) {
                    ans += ways(i, currFuel-used);
                    ans %= MOD;
                }
            }

            return dp[source][currFuel] = ans;
        }

        if (dp[source][currFuel] != -1) {
            return dp[source][currFuel];
        }

        int ans = 0;
        for (int i = 0; i < locs.size(); i++) {
            int used = abs(locs[i]-locs[source]);
            if (i != source &&  used <= currFuel) {
                ans += ways(i, currFuel-used);
                ans %= MOD;
            }
        }
        
        return dp[source][currFuel] = ans;
    }
};

class Solution {
public:
    int dp[100005][27];
    vector<int> times;
    string colors;

    int minCost(string colors, vector<int>& neededTime) {
        memset(dp, -1, sizeof dp); times = neededTime;
        this->colors = colors;
        return findCost(0, 0);
    }

    int findCost(int pos, int lastColor) {
        if (pos == colors.size()) {
            return 0;
        }

        if (dp[pos][lastColor] != -1) {
            return dp[pos][lastColor];
        }

        int mini = 2000000000;
        if (colors[pos]-'a'+1 != lastColor) {
            mini = min(mini, findCost(pos+1, colors[pos]-'a'+1));
        } 

        mini = min(mini, findCost(pos+1, lastColor) + times[pos]);

        return dp[pos][lastColor] = mini;
    }
};

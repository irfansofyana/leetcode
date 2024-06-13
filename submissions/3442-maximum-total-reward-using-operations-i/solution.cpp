class Solution {
public:
    int dp[2005][2005];
    vector<int> rewards;

    int maxTotalReward(vector<int>& rewardValues) {
        memset(dp, -1, sizeof dp);
        rewards = rewardValues;
        sort(rewards.begin(), rewards.end());
        return findMax(0, 0);
    }

    int findMax(int pos, int sum) {
        if (pos == rewards.size()) {
            return sum;
        }

        if (dp[pos][sum] != -1) {
            return dp[pos][sum];
        }

        int maxSum = findMax(pos+1, sum);
        if (rewards[pos] > sum) {
            if (rewards[pos] + sum > 2000) {
                maxSum = max(maxSum, sum+rewards[pos]); 
            } else {
                maxSum = max(maxSum, findMax(pos+1, sum+rewards[pos]));
            }
        }

        return dp[pos][sum] = maxSum;
    }
};

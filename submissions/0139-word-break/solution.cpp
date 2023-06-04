class Solution {
public:
    int dp[300];//s.size()-1-a+1 == idx 
    vector<string> words;

    int findPossible(string s, int idx) {
        if (idx == 0) {
            return 1;
        }
        if (dp[idx] != -1) {
            return  dp[idx];
        }

        string ss = s.substr(s.size()-idx, idx);
        for (int i = 0; i < words.size(); i++) {
            string pref = ss.substr(0, words[i].size());
            if (pref == words[i]) {
                int isPossible = findPossible(s, idx-words[i].size());
                if (isPossible == 1) {
                    return dp[idx] = 1;
                }
            }
        }
        return dp[idx] = 0;
    }

    bool wordBreak(string s, vector<string>& wordDict) {
        memset(dp, -1, sizeof dp);

        words = wordDict;
        int isPossible = findPossible(s, s.size());

        return isPossible == 1 ? true: false;
    }
};

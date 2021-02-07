class Solution {
public:
    int maxScoreWords(vector<string>& words, vector<char>& letters, vector<int>& score) {
        int freq[26] = {0};
        int wordsLength = (int)words.size();
        int wordsScore[15] = {0};

        for (int i = 0; i < letters.size(); ++i) freq[letters[i]-'a']++;
        for (int i = 0; i < wordsLength; ++i) {
            for (int j = 0; j < words[i].size(); ++j) {
                wordsScore[i] += score[words[i][j] - 'a'];
            }
        }
        
        int ans = 0;
        for (int i = 0; i < (1 << wordsLength); ++i) {
            int used[26] = {0};
            vector<int> taken;
            
            for (int j = 0; j < wordsLength; ++j) {
                if ((i&(1<<j)) > 0) {
                    taken.push_back(j);
                    for (int k = 0; k < words[j].size(); ++k) {
                        ++used[words[j][k]-'a'];
                    }        
                }
            }
            
            bool can = true;
            for (int j = 0; j < 26; ++j) {
                if (used[j] > freq[j]) can = false;
            }
            
            if (can) {
                int scoreTmp = 0;
                for (int j = 0; j < (int)taken.size(); ++j) {
                    scoreTmp += wordsScore[taken[j]];
                }
                ans = max(ans, scoreTmp);
            }
        }
        
        return ans;
    }
};

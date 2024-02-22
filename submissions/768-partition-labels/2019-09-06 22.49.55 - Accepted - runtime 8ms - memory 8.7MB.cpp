class Solution {
public:
    vector<int> partitionLabels(string S) {
        vector<int> res;
        int sz = (int)S.size();
        int last[30];
        memset(last, -1, sizeof last);
        for (int i = 0; i < sz; ++i){
            last[S[i]-'a'] = i;
        }
        int idx = last[S[0]-'a'];
        int start = 0;
        while (idx < sz){
            for (int j = start; j <= idx; ++j){
                idx = max(idx, last[S[j]-'a']);
            }
            res.push_back(idx-start+1);
            start = idx+1;
            if (start >= sz) break;
            idx = last[S[start]-'a'];
        }
        return res;
    }
};
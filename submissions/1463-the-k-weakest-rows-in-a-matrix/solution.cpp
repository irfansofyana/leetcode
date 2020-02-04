class Solution {
public:
    vector<int> kWeakestRows(vector<vector<int>>& mat, int k) {
        vector<int> result;
        vector<pair<int,int> > tmp;
        int n = (int)mat.size();
        for (int i = 0; i < n; ++i){
            int m = (int)mat[i].size();
            int soldiers = 0;
            for (int j = 0; j < m && mat[i][j] == 1; ++j){
                ++soldiers;
            }
            tmp.push_back({soldiers, i});
        }
        sort(tmp.begin(), tmp.end());
        for (int i = 0; i < k; ++i){
            result.push_back(tmp[i].second);
        }
        return result;
    }
};

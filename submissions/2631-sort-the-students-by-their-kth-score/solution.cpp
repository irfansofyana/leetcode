class Solution {
public:
    int globalK;

    vector<vector<int>> sortTheStudents(vector<vector<int>>& score, int k) {
        globalK = k;

        auto comp = [this](const vector<int>& a, const vector<int>& b) {
            return a[globalK] > b[globalK];
        };

        sort(score.begin(), score.end(), comp);
        
        return score;
    }
};

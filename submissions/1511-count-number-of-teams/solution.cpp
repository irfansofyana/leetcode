class Solution {
public:
    int numTeams(vector<int>& rating) {
        int cnt = 0;
        int numberOfTeams = (int)rating.size();
        
        for (int i = 0; i < numberOfTeams; ++i) {
            for (int j = i + 1; j < numberOfTeams; ++j) {
                for (int k = j + 1; k < numberOfTeams; ++k){
                    if (rating[i] > rating[j] && rating[j] > rating[k] ||
                       rating[i] < rating[j] && rating[j] < rating[k]) {
                        ++cnt;
                    }
                }
            }
        }
        
        return cnt;
    }
};

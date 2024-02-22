class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> res;
        for (int i = 0; i <= num; ++i){
            if (i == 0) res.push_back(0);
            else res.push_back(res[i/2] + (i%2 == 1));
        }
        return res;
    }
};
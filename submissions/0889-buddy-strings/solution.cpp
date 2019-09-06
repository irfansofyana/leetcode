class Solution {
public:
    bool buddyStrings(string A, string B) {
        vector<int> beda;
        if ((int)A.size() != (int)B.size()) return false;
        else {
            for (int i = 0; i < (int)A.size(); ++i){
                if (A[i] != B[i]) beda.push_back(i);
            }
            int sz = (int)beda.size(); 
            if (sz == 0){
                sort(A.begin(), A.end());
                for (int i = 0; i < (int)A.size(); ++i){
                    if (A[i] == A[i+1]) return true;
                }
                return false;
            }else if (sz == 1) {
                return false;
            }else if (sz == 2){
                  swap(A[beda[0]], A[beda[1]]);
                  return (A == B);
            }else {
                return false;
            }
        }
    }
};

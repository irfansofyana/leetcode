class Solution {
public:
    int countTriplets(vector<int>& arr) {
        int n = (int)arr.size();
        int val[n+1][n+1];
        
        for (int i = 0; i < n; i++) {
            int xorVal = arr[i];
            val[i][i] = xorVal;
            for (int j = i+1; j < n; j++) {
                xorVal ^= arr[j];
                val[i][j] = xorVal;
            }
        }
        
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i+1; j < n; j++) {
                for (int k = j; k < n; k++) {
                    if (val[i][j-1] == val[j][k]) {
                        cnt++;
                    }
                }
            }
        }
        
        return cnt;
    }
};
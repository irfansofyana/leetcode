class Solution {
public:
    string convert(string s, int numRows) {
        int len = (int)s.size();
        vector<char> arr[numRows + 1];
            
        int x = 0;
        bool down = true;
        for (int i = 0; i < len; ++i){
            arr[x].push_back(s[i]);
            if (down) {
                if (x == numRows - 1) {
                    if (x != 0) {
                        down = false;
                        --x;
                    }
                } else {
                    ++x;
                }
            } else {
                if (x == 0) {
                    down = true;
                    ++x;
                } else {
                    --x;
                }
            }
        }
        
        string ans = "";
        for (int i = 0; i < numRows; ++i){
            for (int j = 0; j < arr[i].size(); ++j){
                ans += arr[i][j];
            }
        }
        
        return ans;
    }
};
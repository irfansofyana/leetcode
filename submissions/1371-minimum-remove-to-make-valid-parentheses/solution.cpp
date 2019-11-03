class Solution {
public:
    string minRemoveToMakeValid(string s) {
        int n = (int)s.size();
        string ans = "";
        int ret = 0;
        for (int i = 0; i < n; ++i){
            if (s[i] == '('){
                ++ret;
            }else if (s[i] == ')'){
                --ret;
                if (ret < 0){
                    s[i] = '*';
                    ret = 0;
                }
            }
        }
        for (int i = n-1; i >= 0 && ret > 0; --i){
            if (s[i] == '(') {
                s[i] = '*';
                --ret;
            }
        }
        for (int i = 0; i < (int)s.size(); ++i){
            if (s[i] == '*') continue;
            ans += s[i];
        }
        return ans;
    }
};

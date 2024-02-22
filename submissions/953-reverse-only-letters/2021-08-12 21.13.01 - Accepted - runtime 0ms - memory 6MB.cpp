class Solution {
public:
    string reverseOnlyLetters(string s) {
        vector<char> arc; int n = (int)s.size();
        for (int i = 0; i < n; i++) {
            if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
                arc.push_back(s[i]);
            }
        }
        reverse(arc.begin(), arc.end());
        int idx = 0;
        for (int i = 0; i < n; i++) {
            if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
                s[i] = arc[idx++];
            }
        }
        return s;
    }
};
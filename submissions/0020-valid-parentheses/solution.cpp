class Solution {
public:
    char oppositeChar(char ch) {
        if (ch == '(') return ')';
        else if (ch == '[') return ']';
        else return '}';
    }
    
    bool isValid(string s) {
        stack<char> st;
        for (int i = 0; i < s.size(); ++i) {
            if (s[i] == '(' || s[i] == '[' || s[i] == '{') {
                st.push(s[i]);
            }
            else {
                if (st.empty()) {
                    return false;
                } else if (oppositeChar(st.top()) != s[i]) {
                    return false;
                } else {
                    st.pop();
                }
            }
        }
        return st.empty();
    }
};

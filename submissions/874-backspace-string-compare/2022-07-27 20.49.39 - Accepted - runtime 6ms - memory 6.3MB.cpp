class Solution {
public:
    stack<char> stackS, stackT;
    
    bool backspaceCompare(string s, string t) {
        fillStack(&stackS, s);
        fillStack(&stackT, t);
        
        return isSameStack();
    }
    
    void fillStack(stack<char> *st, string s) {
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '#' && (*st).size() != 0) {
                (*st).pop();
                continue;
            }
            
            if (s[i] != '#') {
                (*st).push(s[i]);
            }
        }
    }
    
    bool isSameStack() {
        if (stackS.size() != stackT.size()) {
            return false;
        }
        
        while (!stackS.empty()) {
            if (stackS.top() != stackT.top()) {
                return false;
            }
            
            stackS.pop();
            stackT.pop();
        }
        
        return true;
    }
};
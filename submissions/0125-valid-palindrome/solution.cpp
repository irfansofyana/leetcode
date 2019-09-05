class Solution {
public:
    string get(string s){
        string front = "";
        for (int i = 0; i < s.size(); ++i){
            if (s[i] >= 'a' && s[i] <= 'z') 
                front += s[i];
            else if (s[i] >= 'A' && s[i] <= 'Z')
                front += (s[i]-'A'+'a');
            else if (s[i] >= '0' && s[i] <= '9')
                front += (s[i]);
        }
        return front;
    }
    bool isPalindrome(string s) {
        string a = get(s);
        reverse(s.begin(), s.end());
        string b = get(s);
        return a == b;
    }
};

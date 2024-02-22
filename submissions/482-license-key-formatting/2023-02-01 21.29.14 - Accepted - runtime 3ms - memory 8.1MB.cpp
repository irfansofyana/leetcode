class Solution {
public:
    string licenseKeyFormatting(string s, int k) {
        string res = "";
        int curr = 0;
        bool foundNonDash = true;

        for (int i = s.size()-1; i >= 0; i--) {
            if (s[i] == '-') {
                continue;
            }

            if (curr == 0 && !foundNonDash) {
                res += '-';
            }

            foundNonDash = false;
            res += toupper(s[i]);

            curr++;
            if (curr == k) {
                curr = 0;
            }
        }

        reverse(res.begin(), res.end());

        return res;
    }
};
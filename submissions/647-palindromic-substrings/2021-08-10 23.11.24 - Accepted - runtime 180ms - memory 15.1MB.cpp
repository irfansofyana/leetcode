class Solution {
public:
    int dp[1002][1002];
    int isPalindrome[1002][1002];
    string st;
    
    int isPalindromeSubstring(int l, int r) {
        if (l == r) {
            return isPalindrome[l][r] = 1;
        } else if (l > r) {
            return isPalindrome[l][r] = 0;
        } else if (isPalindrome[l][r] != -1) {
            return isPalindrome[l][r];
        } 
        
        if (st[l] != st[r]) {
            return isPalindrome[l][r] = 0;
        }
        if (r-l == 1) {
            isPalindrome[l][r] = 1;
        } else {
            isPalindrome[l][r] = isPalindromeSubstring(l+1, r-1);
        }
        return isPalindrome[l][r];
    }
    
    int countPalindromicSubstring(int left, int right) {
        if (left == right) {
            return 1;
        }
        if (left > right) {
            return 0;
        }
        if (dp[left][right] != -1) {
            return dp[left][right];
        }
        
        int cnt = countPalindromicSubstring(left+1, right) + countPalindromicSubstring(left, right-1);
        if (st[left] == st[right] && isPalindromeSubstring(left, right) == 1) {
            cnt += 1;
        }
        cnt -= countPalindromicSubstring(left+1, right-1);
        
        return dp[left][right] = cnt;
    }
    
    int countSubstrings(string s) {
        memset(dp, -1, sizeof dp);
        memset(isPalindrome, -1, sizeof isPalindrome);
        
        int n = s.size();
        st = s;
        return countPalindromicSubstring(0, n-1);
    }
};
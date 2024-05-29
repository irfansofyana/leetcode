class Solution {
public:
    int dp[(1<<15)+1][16];
    int numElements;

    int countArrangement(int n) {
        memset(dp, -1, sizeof dp);    
        numElements = n;
        return compute(0, 0);
    }

    int compute(int bitmask, int pos) {
        if (pos == numElements) {
            return 1;
        }
        if (dp[bitmask][pos] != -1) {
            return dp[bitmask][pos];
        }
        int cnt = 0;
        for (int i = 0; i < numElements; i++) {
            bool isNotUsed = (bitmask&(1<<i)) == 0;
            bool isBeautiful = ((i+1) % (pos+1) == 0) || ((pos+1) % (i+1) == 0);
            if (isNotUsed && isBeautiful) {
                cnt += compute((bitmask|(1<<i)), pos+1);
            }
        }

        return dp[bitmask][pos] = cnt;
    }
};

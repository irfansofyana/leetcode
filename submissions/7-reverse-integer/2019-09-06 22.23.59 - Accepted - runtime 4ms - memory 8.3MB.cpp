class Solution {
public:
    long long reverse(long long x) {
        long long absolut = abs(x);
        long long res = 0LL;
        while (absolut > 0){
            res = (res * 10) + absolut%10;
            absolut /= 10;
        }
        if (res >= (1LL << 31)) return 0;
        if (x < 0) res *= -1;
        return res;
    }
};
class Solution {
public:
    int integerReplacement(long long n) {
        if (n == 1) return 0;
        else if (n == 2) return 1;
        else if (n%2 == 0) return integerReplacement(n/2) + 1;
        else return min(integerReplacement((n-1)/2), integerReplacement((n+1)/2)) + 2;
    }
};

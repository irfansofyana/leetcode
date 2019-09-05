class Solution {
public:
    bool checkPerfectNumber(int num) {
        long long sum = 0LL;
        for (int i = 1; i <= (int)sqrt(num); ++i){
            if (num%i == 0) {
                int j = num / i;
                if (i < num) sum += i;
                if (i != j && j < num) sum += j;
            }
        } 
        return (sum == num && num > 0);
    }
};

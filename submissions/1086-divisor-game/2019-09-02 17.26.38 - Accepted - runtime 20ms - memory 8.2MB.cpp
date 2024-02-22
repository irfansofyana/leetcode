class Solution {
public:
    bool dp[1005];
    void isi(){
        dp[1] = 0;
        for (int i = 2; i <= 1000; ++i){
            dp[i] = 0;
            for (int j = 1; j < i; ++j){
                if (i%j == 0 && dp[i-j] == 0){
                    dp[i] = 1;
                    break;
                }
            }
        }
    }
    bool divisorGame(int N) {
        isi();
        return (dp[N] == 1);
    }
};
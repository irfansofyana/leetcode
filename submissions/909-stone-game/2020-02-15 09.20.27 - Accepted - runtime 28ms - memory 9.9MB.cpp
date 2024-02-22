class Solution {
public:
    int dp[505][505];
    int tumpukan[505];
    int inf = -1000000000;
    
    int cari(int i, int j){
        int giliran = (i + j)%2;
        if (i == j){
            return (giliran%2 == 1 ? tumpukan[i] : -tumpukan[i]);
        }
        if (dp[i][j] != inf) return dp[i][j];
        if (giliran%2 == 1){
            return dp[i][j]=max(cari(i+1, j) + tumpukan[i], cari(i, j-1) + tumpukan[j]);
        }else{
            return dp[i][j]=min(cari(i+1, j)-tumpukan[i], cari(i, j-1)-tumpukan[j]);
        }
    }
    
    bool stoneGame(vector<int>& piles) {
        int n = (int)piles.size();
        for (int i = 0; i < n; ++i){
            for (int j = i; j < n; ++j){
                dp[i][j] = inf;
            }
        }
        for (int i = 0; i < n; ++i){
            tumpukan[i] = piles[i];
        }
        // cout << cari(0, (int)piles.size()-1) << '\n';
        return cari(0, (int)piles.size()-1) > 0;
    }
};
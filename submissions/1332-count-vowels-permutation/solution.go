const MOD = 1000000007

func countVowelPermutation(n int) int {
    dp := make([][5]int, n + 1)
    for i := 0; i < 5; i++ {
        dp[1][i] = 1
    }

    for i := 2; i <= n; i++ {
        dp[i][0] += dp[i-1][1]
        dp[i][1] += dp[i-1][0] + dp[i-1][2]
        dp[i][2] += dp[i-1][0] + dp[i-1][4] + dp[i-1][1] + dp[i-1][3]
        dp[i][3] += dp[i-1][2] + dp[i-1][4]
        dp[i][4] += dp[i-1][0]

        for j := 0; j < 5; j++ {
            dp[i][j] %= MOD
        }
    }

    cnt := 0
    for i := 0; i < 5; i++ {
        cnt += dp[n][i] % MOD
        cnt %= MOD
    }

    return cnt
}

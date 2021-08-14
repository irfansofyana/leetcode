func sumOfDigit(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}

func countBalls(lowLimit int, highLimit int) int {
    freq := map[int]int{}
    
    for i := lowLimit; i <= highLimit; i++ {
        sumDig := sumOfDigit(i)
        freq[sumDig] += 1
    }
    
    ans := 0
    for _, v := range freq {
        if v > ans {
            ans = v
        }
    }
    
    return ans
}

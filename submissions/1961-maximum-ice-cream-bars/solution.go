func maxIceCream(costs []int, coins int) int {
    cntArr := make([]int, 100001)

    for _, cost := range costs {
        cntArr[cost]++
    }

    ans := 0
    for i := 1; i <= 100000; i++ {
        maxCan := coins / i
        onlyHave := cntArr[i]

        if maxCan >= onlyHave {
            ans += onlyHave
            coins -= (onlyHave * i)
        } else {
            ans += maxCan
            coins -= (maxCan * i)
        }
    }

    return ans
}

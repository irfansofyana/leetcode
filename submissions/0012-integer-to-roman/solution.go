func intToRoman(num int) string {
    mapping := map[int]string{
        1: "I",
        4: "IV",
        5: "V",
        9: "IX",
        10: "X",
        40: "XL",
        50: "L",
        90: "XC",
        100: "C",
        400: "CD",
        500: "D",
        900: "CM",
        1000: "M",
    }
    denom := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

    result := ""
    n := len(denom)

    for i := n-1; i >= 0; i-- {
        if denom[i] > num {
            continue
        }

        many := num / denom[i]
        for j := 0; j < many; j++ {
            result += mapping[denom[i]]
        }

        num -= many * denom[i]
    }

    return result
}

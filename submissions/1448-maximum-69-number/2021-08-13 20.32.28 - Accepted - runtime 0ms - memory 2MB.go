import "strconv"

func maximum69Number (num int) int {
    strNum := strconv.Itoa(num)
    ans := ""
    for i := 0; i < len(strNum); i++ {
        if strNum[i] == '6' {
            ans = ans + "9"
            ans = ans + strNum[i+1:]
            break
        }
        ans = ans + string(strNum[i]);
    }
    
    res, _ := strconv.Atoi(ans)
    return res
}
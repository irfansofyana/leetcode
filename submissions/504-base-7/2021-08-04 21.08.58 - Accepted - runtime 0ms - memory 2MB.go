import "strconv"

func convertToBase7(num int) string {
    if num == 0 {
        return "0"
    }
    
    var isPositive bool = true
    if num < 0 {
        isPositive = false
        num = -1 * num
    }
    
    base7 := ""
    for num > 0 {
        base7 = strconv.Itoa(num % 7) + base7
        num /= 7
    }
    
    if !isPositive {
        base7 = "-" + base7
    }
    
    return base7
}
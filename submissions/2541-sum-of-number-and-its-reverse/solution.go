func sumOfNumberAndReverse(num int) bool {
    for i := 0; i <= num; i++ {
        reversedI := reverse(i)
        if num-i == reversedI {
            return true
        }
    }

    return false
}

func reverse(a int) int {
    num := 0

    for a > 0 {
        num = 10 * num + a % 10
        a /= 10
    }

    return num
}

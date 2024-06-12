func sumOfTheDigitsOfHarshadNumber(x int) int {
    sum := sumDigits(x)

    if x % sum == 0 {
        return sum
    }

    return -1
}

func sumDigits(num int) int {
    sum := 0
    
    for num > 0 {
        sum += num % 10
        num /= 10
    }

    return sum
}

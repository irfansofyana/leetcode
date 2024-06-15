func countEven(num int) int {
    count := 0

    for i := 1; i <= num; i++ {
        if sumDigits(i) % 2 == 0 {
            count++
        }
    }

    return count
}

func sumDigits(num int) int {
    sum := 0
    
    for num > 0 {
        sum += num % 10
        num /= 10
    }

    return sum
}

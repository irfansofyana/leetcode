func isHappy(n int) bool {
    slow := n; fast := n;
    for true {
        slow = next(slow)
        fast = next(next(fast))
        if slow == fast {
            break
        }
    }
    if slow == 1 {
        return true
    }

    return false
}

func next(n int) int {
    nxt := 0
    for n > 0 {
        digit := n % 10
        nxt += digit * digit
        n /= 10
    }
    return nxt
}

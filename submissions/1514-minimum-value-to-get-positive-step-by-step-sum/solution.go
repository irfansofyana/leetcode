func minStartValue(nums []int) int {
    sum := 0; mini := 0
    for i, num := range nums {
        sum += num
        if i == 0 {
            mini = sum
            continue
        }

        if sum < mini {
            mini = sum
        }
    }

    if mini >= 0 {
        return 1
    }

    return -1 * mini + 1
}

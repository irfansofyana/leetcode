func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
    tot := int64(0)
    for i := 0; i <= total / cost1; i++ {
        rem := total - cost1 * i
        tot += int64(rem/cost2 + 1)
    }
    return tot
}

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    if k <= numOnes + numZeros {
        return min(k, numOnes)
    }

    k -= (numOnes + numZeros)

    return numOnes-k
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}

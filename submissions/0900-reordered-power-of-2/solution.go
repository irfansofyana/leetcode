func reorderedPowerOf2(n int) bool {
    num := 1
    for num <= 1000000000 {
        if isCanBeReorder(num, n) {
            return true
        }
        num *= 2
    }
    
    return false
}

func isCanBeReorder(a, b int) bool {
    cnt1 := make([]int, 10)
    cnt2 := make([]int, 10)

    for a > 0 {
        cnt1[a%10]++
        a /= 10
    }

    for b > 0 {
        cnt2[b%10]++
        b /= 10
    }

    for i := 0; i < 10; i++ {
        if cnt1[i] != cnt2[i] {
            return false
        }
    }

    return true
}

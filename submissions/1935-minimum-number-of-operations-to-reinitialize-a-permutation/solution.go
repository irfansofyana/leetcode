func reinitializePermutation(n int) int {
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = i;
    }

    cnt := 0
    tmp := make([]int, n)
    for true {
        for i := 0; i < n; i++ {
            if i % 2 == 0 {
                tmp[i] = arr[i / 2]
            } else {
                tmp[i] = arr[n / 2 + (i - 1) / 2]
            }
        }
        for i := 0; i < n; i++ {
            arr[i] = tmp[i];
        }

        cnt++
        if isCheck(arr) {
            break
        }
    }

    return cnt
}

func isCheck(arr []int) bool {
    for i := 0; i < len(arr); i++ {
        if arr[i] != i {
            return false
        }
    }

    return true;
}

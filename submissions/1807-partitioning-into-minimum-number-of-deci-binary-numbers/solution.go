import "strings"

func minPartitions(n string) int {
    maks := 0
    for _, d := range n {
        num, _ :=   strconv.Atoi(string(d))
        if num > maks {
            maks = num
        }
    }   

    return maks
}

import "fmt"

func hammingDistance(x int, y int) int {
    hamming := 0
    for i := 0; i < 32; i++ {
        bx := (x&(1<<i))
        by := (y&(1<<i))
        hamming += dist(bx, by)
    }
    
    return hamming
}

func dist(a, b int) int {
    var va, vb int
    if a > 0 {
        va = 1
    }
    if b > 0 {
        vb = 1
    }
    
    return abs(va, vb)
}

func abs(a, b int) int {
    if a > b {
        return a-b
    }
    return b-a
}

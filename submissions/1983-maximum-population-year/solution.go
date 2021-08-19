func maximumPopulation(logs [][]int) int {
    maxYear := 2050; minYear := 1950
    
    maks := 0; ansYear := maxYear
    for i := maxYear; i >= minYear; i-- {
        cnt := 0
        for j := 0; j < len(logs); j++ {
            if logs[j][0] <= i && i < logs[j][1] {
                cnt += 1
            }
        }
        if cnt >= maks {
            ansYear = i
            maks = cnt
        }
    }
    
    return ansYear
}

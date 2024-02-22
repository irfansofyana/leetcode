func brokenCalc(startValue int, target int) int {
    steps := 0
    for target != startValue {
        if target < startValue {
            target++
        } else if target % 2 == 0 {
            target = target / 2
        } else {
            target++
        }
        
        steps++
    }
    
    return steps
}
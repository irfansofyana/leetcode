func wateringPlants(plants []int, capacity int) int {
    now := capacity
    pos := -1 
    steps := 0
    
    for i, plant := range plants {
        if now >= plant {
            now -= plant
            steps += i - pos
            pos = i
            continue
        }
        
        steps += pos + 1 + i + 1
        now = capacity - plant
        pos = i
    }
    
    return steps
}

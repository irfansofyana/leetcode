func allPaths(graph [][]int, currNode int) [][]int {
    if currNode == len(graph) - 1 {
        tmp := []int{currNode}
        return [][]int{tmp}
    }
    
    paths := make([][]int, 0)
    for _, adjNode := range graph[currNode] {
        pathsExist := allPaths(graph, adjNode)
        for _, path := range pathsExist {
            tmp := []int{currNode}
            tmp = append(tmp, path...)
            paths = append(paths, tmp)
        }
    }
    
    return paths
}

func allPathsSourceTarget(graph [][]int) [][]int {
    return allPaths(graph, 0)
}
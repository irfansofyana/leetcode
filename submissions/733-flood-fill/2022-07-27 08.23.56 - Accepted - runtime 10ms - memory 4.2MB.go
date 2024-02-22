func isValidRow(numRow int, numCol int, r int, c int) bool {
    return r >= 0 && r < numRow && c >= 0 && c < numCol    
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    recursive(&image, sr, sc)
    for i := 0; i < len(image); i++ {
        for j := 0; j < len(image[0]); j++ {
            if image[i][j] == -1 {
                image[i][j] = color
            }
        }
    }
    return image
}

func recursive(image *[][]int, r int, c int) {
    oldColor := (*image)[r][c]
    numRow := len(*image)
    numCol := len((*image)[0])
    
    dx := []int{-1, 0, 0, 1}
    dy := []int{0, -1, 1, 0}
    
    (*image)[r][c] = -1
    for i := 0; i < 4; i++ {
        nr := r + dx[i]
        nc := c + dy[i]
        if isValidRow(numRow, numCol, nr, nc) && (*image)[nr][nc] == oldColor {
            recursive(image, nr, nc)
        }
    }
}
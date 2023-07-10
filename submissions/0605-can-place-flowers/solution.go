func canPlaceFlowers(flowerbed []int, n int) bool {
    i := 0
    for i < len(flowerbed) && n > 0 {
       if flowerbed[i] == 0 && i - 1 >= 0 && flowerbed[i-1] == 0 && i + 1 < len(flowerbed) && flowerbed[i+1] == 0 {
           flowerbed[i] = 1
           n--
           i += 2
           continue
       }
       if i == 0 && flowerbed[i] == 0 && i + 1 < len(flowerbed) && flowerbed[i+1] == 0 {
           flowerbed[i] = 1
           n--
           i += 2
           continue
       }
       if i == len(flowerbed) - 1 && i - 1 >= 0 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
           flowerbed[i] = 1
           n--
           i += 2
           continue
       }
       if i == 0 && len(flowerbed) == 1 && flowerbed[0] == 0 {
           n--
           i++
       }
       i++
    }

    return n == 0
}

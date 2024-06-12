func sumOfThree(num int64) []int64 {
 //x + x+1 + x+2 = 3(x+1) = 33
 if num % 3 != 0 {
    return []int64{}
 }   

 x := (num / 3) - 1

 return []int64{x, x+1, x+2}
}

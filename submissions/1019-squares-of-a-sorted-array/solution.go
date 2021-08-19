import "sort"

func sortedSquares(nums []int) []int {
    squares := make([]int, 0)
    for _, num := range nums {
        squares = append(squares, num * num)
    }
    sort.Ints(squares)
    return squares
}

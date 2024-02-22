func containsDuplicate(nums []int) bool {
    var isExist = map[int]bool{}
    for _, num := range(nums) {
        if isExist[num] {
            return true
        }
        isExist[num] = true
    }
    return false
}
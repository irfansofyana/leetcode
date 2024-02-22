import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)

    pairs := make([]int, 0)
    for _, spell := range spells {
        target := success/int64(spell)
        if success % int64(spell) != 0 {
            target++
        }

        idx := findFirstGTE(potions, int(target))
        many := 0
        if int64(potions[idx]) >= target {
            many = len(potions)-idx
        }

        pairs = append(pairs, many)
    }    

    return pairs
}

func findFirstGTE(array []int, target int) int {
    left := 0; right := len(array) - 1
    
    idx := len(array) - 1
    for left <= right {
        mid := (left + right) / 2
        if array[mid] >= target {
            idx = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return idx
}
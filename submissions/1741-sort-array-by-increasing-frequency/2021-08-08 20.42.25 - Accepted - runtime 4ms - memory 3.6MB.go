import "sort"

type pair struct {
    num int
    freq int
}

func frequencySort(nums []int) []int {
    freq := map[int]int{};
    
    for _, num := range nums {
        freq[num] += 1
    }
    
    tmp := make([]pair, 0)
    for k, v := range freq {
        tmp = append(tmp, pair{num: k, freq: v})
    }
    
    sort.SliceStable(tmp, func(i, j int) bool {
        if tmp[i].freq == tmp[j].freq {
            return tmp[i].num > tmp[j].num
        }
        return tmp[i].freq < tmp[j].freq
    })
    
    ans := make([]int, 0)
    for i := 0; i < len(tmp); i++ {
        for j := 0; j < tmp[i].freq; j++ {
            ans = append(ans, tmp[i].num)
        }
    }
    
    return ans
}
import "sort"

type box struct {
    BoxesNumber int
    UnitNumber int
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
    arr := make([]box, 0); n := len(boxTypes)
    for i := 0; i < n; i++ {
        arr = append(arr, box{
            BoxesNumber: boxTypes[i][0],
            UnitNumber: boxTypes[i][1],
        })
    }
    sort.SliceStable(arr, func (i, j int) bool {
        if arr[i].UnitNumber == arr[j].UnitNumber {
            return arr[i].BoxesNumber < arr[j].BoxesNumber
        }
        return arr[i].UnitNumber > arr[j].UnitNumber
    })
    
    ans := 0
    for i := 0; i < n && truckSize > 0; i++ {
        var mini int = arr[i].BoxesNumber
        if truckSize < arr[i].BoxesNumber {
            mini = truckSize
        }
        ans += mini * arr[i].UnitNumber
        truckSize -= mini
    }
    
    return ans
}
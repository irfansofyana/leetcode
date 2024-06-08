/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    mapIdToEmployee := getMappingIdToEmployee(employees)
    var ans int
    for _, e := range employees {
        if e.Id == id {
            ans = getSum(e, mapIdToEmployee)
            break
        }
    }
    return ans
}

func getMappingIdToEmployee(employees []*Employee) map[int]*Employee {
    mapping := make(map[int]*Employee)
    for _, e := range employees {
        mapping[e.Id] = e
    }
    return mapping
}

func getSum(employee *Employee, mapIdToEmployee map[int]*Employee) int {
    sum := employee.Importance
    
    for _, id := range employee.Subordinates {
        e := mapIdToEmployee[id]
        sum += getSum(e, mapIdToEmployee)
    }

    return sum
}

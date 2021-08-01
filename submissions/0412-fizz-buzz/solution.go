import "fmt"

func fizzBuzz(n int) []string {
    result := make([]string, 0)
    
    for i := 1; i <= n; i++ {
        var r string
        if i % 15 == 0 {        
            r = "FizzBuzz"
        } else if i % 3 == 0 {
            r = "Fizz"
        } else if i % 5 == 0  {
            r = "Buzz"
        } else {
            r = fmt.Sprintf("%d", i)
        }
        result = append(result, r)
    }    
    
    return result
}

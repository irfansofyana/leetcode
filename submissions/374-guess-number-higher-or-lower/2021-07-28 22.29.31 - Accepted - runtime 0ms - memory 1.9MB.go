/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l := 1; r := n
    ans := -1
    
    for l <= r {
        mid := (l + r) / 2
        if guess(mid) == 0 {
            ans = mid
            break
        } else if guess(mid) == -1 {
           r = mid - 1
        } else {
            l = mid + 1
        }
    }
    
    return ans
}
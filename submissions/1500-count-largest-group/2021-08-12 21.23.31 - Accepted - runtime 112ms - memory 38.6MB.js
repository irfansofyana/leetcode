/**
 * @param {number} n
 * @return {number}
 */

var sumDigit = function(n) {
    let sum = 0;
    while (n > 0) {
        sum += n % 10;
        n = Math.trunc(n/10);
    }
    return sum;
}

var countLargestGroup = function(n) {
    let freq = {};
    let maks = 0;
    
    for (let i = 1; i <= n; i++) {
        sd = sumDigit(i);
        if (sd in freq) {
            freq[sd] += 1;
        } else {
            freq[sd] = 0;
        }
        if (freq[sd] > maks) {
            maks = freq[sd];
        }
    }
    
    let ans = 0;
    for (let key in freq) {
        if (freq[key] == maks) {
            ans += 1;
        }
    }
    
    return ans;
};
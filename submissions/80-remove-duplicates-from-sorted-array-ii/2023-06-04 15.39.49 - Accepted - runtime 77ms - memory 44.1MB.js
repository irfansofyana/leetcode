/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let curr = 0;
    let idx = 0;

    while (curr < nums.length) {
        nums[idx] = nums[curr];
        idx++;

        let nxt = curr + 1;
        while (nxt < nums.length && nums[nxt] == nums[curr]) {
            if (nxt - curr == 1) {
                nums[idx] = nums[curr];
                idx++;
            }
            nxt++;
        }

        curr = nxt;
    }

    nums = nums.slice(0, idx);

    return idx;
};
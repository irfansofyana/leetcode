class Solution {
public:
    bool isGoodArray(vector<int>& nums) {
        int gcd;
        for (int i = 0; i < (int)nums.size(); ++i){
            gcd = (i == 0 ? nums[i] : __gcd(gcd, nums[i]));
        }
        return (gcd == 1);
    }
};
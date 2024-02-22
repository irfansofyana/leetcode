class Solution {
public:
    vector<long long> maximumEvenSplit(long long finalSum) {
        vector<long long> ans;
        
        if (finalSum % 2 == 1) {
            return ans;
        }
        
        long long sqrtSum = (long long)(sqrt(finalSum));
        while (sqrtSum*(sqrtSum+1) > finalSum) {
            sqrtSum -= 1;
        }
        
        long long sum = 0;
        for (int i = 1; i <= sqrtSum-1; i++) {
            sum += (long long)2*i;
            ans.push_back((long long)2*i);
        }
        ans.push_back(finalSum-sum);
        
        return ans;
    }
};
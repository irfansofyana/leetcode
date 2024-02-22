class Solution {
public:
    bool checkPossibility(vector<int>& nums) {
        int i = 0;
        vector<pair<int,int> > tmp;
        while (i < (int)nums.size()){
            int j = i + 1;
            while (j < (int)nums.size() && nums[j] >= nums[j-1]){
                ++j;
            }
            tmp.push_back({i, j-1});
            i = j;
        }
        if ((int)tmp.size() > 2) return false;
        else if ((int)tmp.size() == 1) return true;
        else {
            bool yes = true;
            int lastf = nums[tmp[0].second];
            for (int i = tmp[1].first+1; i <= tmp[1].second; ++i){
                if (nums[i] < lastf) yes = false; 
            }
            if (yes) return true;
            yes = true;
            int firsts = nums[tmp[1].first];    
            for (int i = tmp[0].second-1; i >= tmp[0].first; --i){
                if (nums[i] > firsts) yes = false;
            }
            return yes;
        }
    }
};
class Solution {
public:
    int minSetSize(vector<int>& arr) {
        int n = (int)arr.size();
        vector<pair<int,int> > new_arr;
        map<int,int> frek;
        
        for (int i = 0; i < n; ++i){
            frek[arr[i]]++;
        }
        for (auto a : frek){
            new_arr.push_back({a.second, a.first});
        }
        sort(new_arr.begin(), new_arr.end());
        int ans = 0;
        int sum = 0;
        for (int i = (int)new_arr.size()-1; i >= 0; --i){
            sum += new_arr[i].first;
            ++ans;
            if (sum >= n/2) {
                break;
            }
        }
        return ans;
    }
};

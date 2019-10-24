class Solution {
public:
    bool comp(pair<int,pair<int,int>> a,pair<int,pair<int,int>> b){
        if (a.second.first == b.second.first) return a.second.second > b.second.second;
        return a.second.first < b.second.first;
    }
    
    int jobScheduling(vector<int>& startTime, vector<int>& endTime, vector<int>& profit) {
        vector<pair<int,pair<int,int> > > arr;
        for (int i = 0; i < (int)startTime.size(); ++i){
            arr.push_back(make_pair(endTime[i],make_pair(startTime[i], profit[i])));
        }
        sort(arr.begin(), arr.end());
        int n = (int)startTime.size();
        int *table = new int[n];
        table[0] = arr[0].second.second;
        for (int i = 1; i < n; ++i){
            int tprofit = arr[i].second.second;
            int lo = 0;
            int hi = i-1;
            int ret = -1;
            while (lo <= hi){
                int mid = (lo + hi) >> 1;
                if (arr[mid].first <= arr[i].second.first){
                    ret = mid;
                    lo = mid+1;
                }else {
                    hi = mid - 1;
                }
            }
            if (ret != -1){
                tprofit += table[ret];
            }
            table[i] = max(tprofit, table[i-1]);
        }
        return table[n-1];
    }
};

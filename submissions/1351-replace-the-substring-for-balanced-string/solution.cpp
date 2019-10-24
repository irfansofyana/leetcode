class Solution {
public:
    int arr[4] = {0, 0, 0, 0};
    
    int indeks(char c){
        if (c == 'Q') return 0;
        else if (c == 'W') return 1;
        else if (c == 'E') return 2;
        else return 3;
    }
    
    bool can(int x, string s){
        int tmp[4] = {0, 0, 0, 0};
        int must = (int)s.size() / 4;
        for (int i = 0; i < (int)s.size(); ++i){
            if (i < x){
                ++tmp[indeks(s[i])];
            }else {
                int sum = 0;
                bool cek = true;
                for (int j = 0; j < 4; ++j){
                    sum += (must-(arr[j] - tmp[j]));
                    if (arr[j]-tmp[j] > must){
                        cek = false;
                    }
                }
                if (cek && sum == x) return true;
                --tmp[indeks(s[i-x])];
                ++tmp[indeks(s[i])];
            }
        }
        int sum = 0;
        bool cek = true;
        for (int j = 0; j < 4; ++j){
            sum += (must-(arr[j] - tmp[j]));
            if (arr[j]-tmp[j] > must){
                cek = false;
            }
        }
        if (cek && sum == x) return true;
        return false;
    }
    
    int balancedString(string s) {
        //0 -> q, 1 -> w, 2 -> e, 3 -> r
        int must = (int)s.size() / 4;
        for (int i = 0; i < (int)s.size(); ++i){
            ++arr[indeks(s[i])];
        }
        int ans = (int)s.size();
        int lo = 0;
        int hi = (int)s.size();
        while (lo <= hi){
            int mid = (lo + hi) >> 1;
            if (can(mid, s)){
                ans = mid;
                hi = mid-1;
            }else {
                lo = mid + 1;
            }
        }
        return ans;
    }
};

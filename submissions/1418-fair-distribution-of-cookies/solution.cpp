class Solution {
public:
    vector<int> arr;
    int n, mini;
    int choosen[10];

    int distributeCookies(vector<int>& cookies, int k) {
        n = cookies.size();   
        arr = cookies;
        mini = 1000000000;
        recursive(0, k); 
        return mini;
    }

    void recursive(int curr, int k) {
        if (curr == n) {
            int mep[k+1];

            for (int i = 0; i < k; i++) {
                mep[i] = 0;
            }
            for (int i = 0; i < n; i++) {
                mep[choosen[i]] += arr[i];
            }
            int localMax = -1;
            for (int i = 0; i < k; i++) {
                localMax = max(localMax, mep[i]);
            }
            mini = min(mini, localMax);

            return;
        }

        for (int i = 0; i < k; i++) {
            choosen[curr] = i;
            recursive(curr + 1, k);
        }
    }
};

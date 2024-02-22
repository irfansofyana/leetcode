class Solution {
public:
    int numRescueBoats(vector<int>& people, int limit) {
        sort(people.begin(), people.end());
        stack<int> st;
        
        int boats = 0;
        for (int i = (int)people.size()-1; i >= 0; i--) {
            if (!st.empty() && st.top() >= people[i]) {
                st.pop();
                continue;
            }
            boats++;
            if (limit-people[i] > 0) {
                st.push(limit-people[i]);
            }
        }
        
        return boats;
    }
};
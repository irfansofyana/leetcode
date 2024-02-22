class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        map<string,vector<string> > mep;
        
        for (auto str: strs) {
            string cstr = str;
            sort(cstr.begin(), cstr.end());
            mep[cstr].push_back(str);
        }
        
        vector<vector<string> > ans;
        
        for (auto k: mep) {
            ans.push_back(k.second);
        }
        
        return ans;
    }
};
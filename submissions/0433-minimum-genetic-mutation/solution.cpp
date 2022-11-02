class Solution {
public:
    int minMutation(string start, string end, vector<string>& bank) {
        queue<string> q;
        map<string, int> numMutations;
        char validGene[] = {'A', 'C', 'G', 'T'};
        
        map<string, bool> isValidGenes;
        for (int i = 0; i < bank.size(); i++) {
            isValidGenes[bank[i]] = true;
        }
        
        q.push(start);
        numMutations[start] = 0;
        while (!q.empty()) {
            string now = q.front();
            q.pop();
            
            for (int i = 0; i < now.size(); i++) {
                string pref = now.substr(0, i);
                string suff = now.substr(i+1, (int)now.size()-1-i);
                
                for (int j = 0; j < 4; j++) {
                    if (now[i] != validGene[j]) {
                        string formed = pref + validGene[j] + suff;
                        if (!isValidGenes[formed] || isValidGenes.find(formed) == isValidGenes.end()) {
                            continue;
                        }
                        
                        if (numMutations.find(formed) == numMutations.end()) {
                            numMutations[formed] = numMutations[now] + 1;
                            q.push(formed);
                        }
                        else if (numMutations[now] + 1 < numMutations[formed]) {
                            numMutations[formed] = numMutations[now] + 1;
                            q.push(formed);
                        }
                    }        
                }
            }
        }
        
        return numMutations.find(end) == numMutations.end() ? -1 : numMutations[end];
    }
};

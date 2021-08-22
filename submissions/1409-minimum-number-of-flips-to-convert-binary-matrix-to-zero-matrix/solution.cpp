class Solution {
public:
    map<vector<vector<int>>, int> dist;
    int dx[4] = {-1, 0, 0, 1};
    int dy[4] = {0, -1, 1, 0};
    queue<vector<vector<int>>> q;
        
    vector<vector<int>> toggle(vector<vector<int>> state, int r, int c) {
        state[r][c] = 1-state[r][c];
        int n = (int)state.size();
        int m = (int)state[0].size();
        
        for (int i = 0; i < 4; i++) {
            if (r + dx[i] >= 0 && r + dx[i] < n && c + dy[i] >= 0 && c + dy[i] < m) {
                state[r + dx[i]][c + dy[i]] = 1 - state[r + dx[i]][c + dy[i]];
            }
        }
        return state;
    }
    
    int minFlips(vector<vector<int>>& mat) {
        int n = (int)mat.size();
        int m = (int)mat[0].size();
        
        dist[mat] = 0;
        q.push(mat);
        
        while (!q.empty()) {
            vector<vector<int>> state = q.front();
            q.pop();
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < m; j++) {
                    vector<vector<int>> nextState = toggle(state, i, j);
                    if ((dist.find(nextState) == dist.end()) ||  (dist[state] + 1 < dist[nextState])) {
                        dist[nextState] = dist[state] + 1;
                        q.push(nextState);
                    }
                }
            }
        }
        
        vector<vector<int>> zero;
        for (int i = 0; i < n; i++) {
            vector<int> tmp;
            for (int j = 0; j < m; j++) {
                tmp.push_back(0);
            }
            zero.push_back(tmp);
        }
        
        if (dist.find(zero) == dist.end()) {
            return -1;
        }
        return dist[zero];
    }
};

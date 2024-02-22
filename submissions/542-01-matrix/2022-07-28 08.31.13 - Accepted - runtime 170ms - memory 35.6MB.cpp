class Solution {
public:
    int n, m;
    int dx[4] = {-1, 0, 0, 1};
    int dy[4] = {0, -1, 1, 0};
    
    bool valid(int r, int c) {
        return (r >= 0 && r < n && c >= 0 && c < m);    
    }
    
    vector<vector<int>> updateMatrix(vector<vector<int>>& mat) {
        vector<vector<int>> dist;
        queue<pair<int,int>> que;
        
        n = mat.size(); m = mat[0].size();
        for (int i = 0; i < n; i++) {
            vector<int> tmp;
            for (int j = 0; j < m; j++) {
                if (mat[i][j] == 0) {
                    tmp.push_back(0);
                    que.push({i, j});
                } else {
                    tmp.push_back(INT_MAX);
                }
            }
            dist.push_back(tmp);
        }
        
        while(!que.empty()) {
            pair<int,int> p = que.front();
            que.pop();
            
            for (int i = 0; i < 4; i++) {
                int px = p.first + dx[i];
                int py = p.second + dy[i];
                
                if (valid(px, py) && dist[p.first][p.second] + 1 < dist[px][py]) {
                    dist[px][py] = dist[p.first][p.second] + 1;
                    que.push({px, py});
                }
            }
        }
        
        return dist;
    }
};
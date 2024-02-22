class Solution {
public:
    int dx[4] = {-1, 0, 0, 1};
    int dy[4] = {0, -1, 1, 0};
    int n, m;
    
    bool isInsideGrid(int r, int c) {
        return (r >= 0 && r < n && c >= 0 && c < m);
    }
    
    int orangesRotting(vector<vector<int>>& grid) {
        vector<vector<int> > dist;
        n = (int)grid.size(); m = (int)grid[0].size();
        queue<pair<int,int> > q;
        
        for (int i = 0; i < n; i++) {
            vector<int> tmp;
            for (int j = 0; j < m; j++) {
                if (grid[i][j] == 0 || grid[i][j] == 1) {
                    tmp.push_back(INT_MAX);
                } else if (grid[i][j] == 2) {
                    tmp.push_back(0);
                    q.push({i, j});
                }
            }
            dist.push_back(tmp);
        }
        
        while (!q.empty()) {
            pair<int,int> now = q.front(); q.pop();
            for (int i = 0; i < 4; i++) {
                int nx = now.first + dx[i];
                int ny = now.second + dy[i];
                
                if (isInsideGrid(nx, ny) && grid[nx][ny] == 1 && dist[now.first][now.second] + 1 < dist[nx][ny]) {
                    dist[nx][ny] = dist[now.first][now.second] + 1;
                    grid[nx][ny] = 2;
                    q.push({nx, ny});
                }
            }
        }
        
        int minimumTime = INT_MIN;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (grid[i][j] == 0) {
                    continue;
                }
                
                if (grid[i][j] == 2) {
                    minimumTime = max(minimumTime, dist[i][j]);
                    continue;
                }
                
                if (grid[i][j] == 1) {
                    return -1;   
                }
            }
        }
        
        return (minimumTime == INT_MIN ? 0: minimumTime);
    }
};
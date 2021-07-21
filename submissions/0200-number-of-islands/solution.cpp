class Solution {
public:
    bool isInsideGrid(int m, int n, int x, int y){
        return x>=0 && x < m && y >= 0 && y < n;    
    }
    
    int numIslands(vector<vector<char>>& grid) {
        int m = (int)grid.size();
        int n = (int)grid[0].size();
        int ans = 0;
        int dx[4] = {-1, 0, 0, 1};
        int dy[4] = {0, -1, 1, 0};
        queue<pair<int,int> > Que;
        
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == '1') {
                    ++ans;
                    Que.push({i, j});
                    grid[i][j] = '0';
                    
                    while (!Que.empty()) {
                        int x = Que.front().first;
                        int y = Que.front().second;
                        Que.pop();
                        for (int k = 0; k < 4; ++k) {
                            int nx = x + dx[k];
                            int ny = y + dy[k];
                            if (isInsideGrid(m, n, nx, ny)) {
                                if (grid[nx][ny] == '1') {
                                    Que.push({nx, ny});
                                    grid[nx][ny] = '0';
                                }
                            }
                        }
                    }
                }
            }
        }
        
        return ans;
    }
};

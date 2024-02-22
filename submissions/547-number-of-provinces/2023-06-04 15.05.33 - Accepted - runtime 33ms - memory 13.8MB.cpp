class Solution {
public:
    int parent[205];

    int findParent(int x) {
        if (x == parent[x]) {
            return x;
        }

        return parent[x] = findParent(parent[x]);
    }

    void merge(int a, int b) {
        int pa = findParent(a);
        int pb = findParent(b);
        parent[pa] = pb;
    }

    int findCircleNum(vector<vector<int>>& isConnected) {
        int n = (int)isConnected.size();

        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }

        for (int i = 0; i < n; i++) {
            for (int j = i+1; j < n; j++) {
                if (isConnected[i][j]) {
                    merge(i, j);
                }
            }
        }

        set<int> provinces;

        for (int i = 0; i < n; i++) {
            provinces.insert(findParent(i));
        }

        return provinces.size();
    }
};
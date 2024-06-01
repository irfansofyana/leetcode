class Solution {
public:
    int parent[200005];

    bool validPath(int n, vector<vector<int>>& edges, int source, int destination) {
        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }    

        for (auto edge: edges) {
            int u = edge[0];
            int v = edge[1];
            connect(u, v);
        }

        return findParent(source) == findParent(destination);
    }

    int findParent(int u) {
        if (u == parent[u]) {
            return u;
        }

        return parent[u] = findParent(parent[u]);
    }

    void connect(int u, int v) {
        int pu = findParent(u);
        int pv = findParent(v);
        if (pu != pv) {
            parent[pu] = pv;
        }
    }
};

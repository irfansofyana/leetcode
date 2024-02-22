int distance(vector<int> a) {
    int dist = 0;
    for (int i = 0; i < a.size(); i++) {
        dist += a[i] * a[i];
    }
    return dist;
}

bool comp(vector<int> a, vector<int> b) {
    return distance(a) < distance(b);
}

class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        sort(points.begin(), points.end(), comp);
        vector<vector<int> > nearest;
        for (int i = 0; i < k; i++) {
            nearest.push_back(points[i]);
        }
        return nearest;
    }
};
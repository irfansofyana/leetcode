class Solution {
public:
    bool checkStraightLine(vector<vector<int>>& coordinates) {
        vector<pair<int,int> > points;
        for (int i = 0; i < (int)coordinates.size(); ++i){
            points.push_back(make_pair(coordinates[i][0], coordinates[i][1]));
        }
        sort(points.begin(), points.end());
        double grad;
        for (int i = 0; i < (int)points.size()-1; ++i){
            int dy = points[i+1].second-points[i].second;
            int dx = points[i+1].first - points[i].first;
            if (dx == 0) return false;
            else if (i == 0) grad = (1.00 * dy) / dx;
            else if (1.00 * dy / dx != grad) return false;
        }
        return true;
    }
};
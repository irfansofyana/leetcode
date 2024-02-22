class Solution {
public:
    int minimumSwap(string s1, string s2) {
        int tx = 0;
        int ty = 0;
        for (int i = 0; i < s1.size(); ++i){
            if (s1[i] == 'x') ++tx;
            else ++ty;
        }
        for (int i = 0; i < s2.size(); ++i){
            if (s2[i] == 'x') ++tx;
            else ++ty;
        }
        if (tx%2 == 1 || ty%2 == 1) {
            return -1;
        }else {
            int xy = 0;
            int yx = 0;
            for (int i = 0; i < s1.size(); ++i){
                if (s1[i] == 'x' && s2[i] == 'y') ++xy;
                else if (s1[i] == 'y' && s2[i] == 'x') ++yx;
            }
            return (xy/2 + yx/2 + (xy%2 == 1 && yx%2 == 1 ? 2 : 0));
        }
    }
};



/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        queue<pair<TreeNode*, int>> q;
        map<int,vector<int>> mep;
        
        q.push(make_pair(root, 0));
        int depth = -1;
        while (!q.empty()) {
            pair<TreeNode*,int> f = q.front();
            q.pop();
            
            if (f.first == NULL) {
                break;
            }
            
            mep[f.second].push_back(f.first->val);
            depth = max(depth, f.second);
            
            if (f.first->left != NULL) {
                q.push(make_pair(f.first->left, f.second+1));
            }
            if (f.first->right != NULL) {
                q.push(make_pair(f.first->right, f.second+1));
            }
        }

        vector<vector<int>> ans;
        for (int i = 0; i <= depth; i++) {
            vector<int> arr = mep[i];
            if (i % 2 == 1) {
                reverse(arr.begin(), arr.end());
            }
            ans.push_back(arr);
        }

        return ans;
    }
};
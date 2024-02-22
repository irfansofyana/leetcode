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
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int> > ans;
        if (root == NULL) {
            return ans;
        }
        
        map<int,vector<int> > mep;
        queue<pair<int,TreeNode*> > q;
        q.push({0, root});
        
        while (!q.empty()) {
            pair<int,TreeNode*> el = q.front();
            q.pop();
            mep[el.first].push_back(el.second->val);
            if (el.second -> left != NULL)
                q.push({el.first + 1, el.second->left});
            if (el.second -> right != NULL)
                q.push({el.first + 1, el.second->right});
        }
        
        for (auto it = mep.begin(); it != mep.end(); it++) {
            ans.push_back(it->second);
        }
        return ans;
    }
};
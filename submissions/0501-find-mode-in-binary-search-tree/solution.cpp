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
    map<int,int> mep;

    void traverse(TreeNode* root) {
        if (root == NULL) {
            return;
        }

        mep[root->val]++;
        if (root->left != NULL) {
            traverse(root->left);
        }
        if (root->right != NULL) {
            traverse(root->right);
        }
    }

    vector<int> findMode(TreeNode* root) {
        traverse(root);
        int maks = -1;
        vector<int> ans;
        for (auto a: mep) {
            int num = a.first;
            int freq = a.second;
            if (freq > maks) {
                maks = freq;
                ans.clear();
                ans.push_back(num);
            } else if (freq == maks) {
                ans.push_back(num);
            }
        }

        return ans;
    }
};
